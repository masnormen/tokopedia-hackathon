package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/masnormen/tokopedia-hackathon/delivery"
	"github.com/masnormen/tokopedia-hackathon/repository/pgsql"
	"github.com/masnormen/tokopedia-hackathon/usecase"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"os"
)

func connectDB() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3307)/tco?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func homePage(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("ok"))
	return
}

func migration(db *gorm.DB) error {
	fmt.Println("Doing migration...")
	err := db.AutoMigrate(
		&pgsql.Buyer{},
		&pgsql.Courier{},
		&pgsql.CourierCostMapping{},
		&pgsql.Product{},
		&pgsql.Seller{},
	)

	if err != nil {
		panic(err)
	}

	return nil
}

func main() {
	db := connectDB()

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "migrate":
			migration(db)
			break
		}
	}

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/health", homePage).Methods("GET")

	// Init
	productOrm := pgsql.NewProductOrm(db)
	searchPageUsecase := usecase.NewSearchPageUsecase(productOrm)
	delivery.NewSearchPageHandler(r, searchPageUsecase)

	fmt.Print("Running...")
	_ = http.ListenAndServe(":9090", r)
}
