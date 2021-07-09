package main

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/gorilla/mux"
	"github.com/masnormen/tokopedia-hackathon/delivery"
	"github.com/masnormen/tokopedia-hackathon/repository/pgsql"
	"github.com/masnormen/tokopedia-hackathon/usecase"
	"net/http"
	"os"
)

func connectDB() *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "user",
		Password: "secret",
		Database: "tco",
	})
	return db
}

func homePage(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("ok"))
	return
}

func migration(db *pg.DB) error {
	models := []interface{}{
		(*pgsql.Buyer)(nil),
		(*pgsql.Courier)(nil),
		(*pgsql.CourierCostMapping)(nil),
		(*pgsql.Product)(nil),
		(*pgsql.Seller)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
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

	_ = http.ListenAndServe(":9090", r)
}
