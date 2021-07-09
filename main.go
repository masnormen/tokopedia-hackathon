package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/masnormen/tokopedia-hackathon/delivery"
	"github.com/masnormen/tokopedia-hackathon/fixtures"
	mysqlRepository "github.com/masnormen/tokopedia-hackathon/repository/mysql"
	"github.com/masnormen/tokopedia-hackathon/usecase"
	"github.com/rs/cors"
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

func main() {
	db := connectDB()

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "migrate":
			err := fixtures.Migrate(db)
			panic(err)
			break
		case "seed":
			fixtures.Seed(db)
			break
		}
		return
	}

	r := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})
	r.HandleFunc("/api/v1/health", homePage).Methods("GET")

	// Init
	productOrm := mysqlRepository.NewProductOrm(db)
	searchPageUsecase := usecase.NewSearchPageUsecase(productOrm)
	delivery.NewSearchPageHandler(r, searchPageUsecase)

	fmt.Print("Running...")
	handler := c.Handler(r)
	_ = http.ListenAndServe(":9090", handler)
}
