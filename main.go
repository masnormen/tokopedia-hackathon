package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/Pallinder/go-randomdata"
	"github.com/gorilla/mux"
	"github.com/masnormen/tokopedia-hackathon/delivery"
	"github.com/masnormen/tokopedia-hackathon/repository/pgsql"
	"github.com/masnormen/tokopedia-hackathon/usecase"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	fmt.Println("Doing migration...")

	return nil
}

func seed(db *gorm.DB) {
	for i := 1; i < 10; i++ {
		index := strconv.Itoa(i)
		products := make([]*pgsql.Product, 0)
		for j := 1; j < 10; j++ {
			product := &pgsql.Product{
				Name:            "Sepatu " + strconv.Itoa(i*j),
				Price:           rand.Intn(350000),
				Rating:          3.5,
				TotalSales:      rand.Intn(1200),
				Weight:          1,
				ProductImageURL: "https://cdn-3.tstatic.net/jualbeli/img/2020/7/2353564/1-26055081-Sepatu-Converse-CT-I-ALL-Star-OX-Black-Original-BNIB-Vietnam---Madiun.jpg",
				SellerID:        i,
			}
			products = append(products, product)
		}
		seller := &pgsql.Seller{
			ID:              i,
			Name:            randomdata.SillyName(),
			Province:        "Jawa Barat",
			City:            "Bekasi",
			Address:         "Jl Candrabaga " + index,
			Postcode:        randomdata.PostalCode("SE"),
			Latitude:        "",
			Longitude:       "",
			ProfileImageURL: "https://pbs.twimg.com/profile_images/1407698877914902530/Uy5uB6Qb_400x400.jpg",
			Badge:           "https://pbs.twimg.com/media/E2txKPEUcAEyjMB.jpg",
			Product:         products,
		}
		courierCostMappings := make([]*pgsql.CourierCostMapping, 0)
		for j := 0; j < 10; j++ {
			CourierCostMapping := &pgsql.CourierCostMapping{
				ID:                      j,
				CourierID:               i,
				ServiceName:             randomdata.SillyName(),
				PostcodeCitySource:      randomdata.PostalCode("SE"),
				PostcodeCityDestination: randomdata.PostalCode("SE"),
				BasePrice:               randomdata.Number(10000, 50000),
				CoefWeight:              randomdata.Decimal(10000, 20000),
			}
			courierCostMappings = append(courierCostMappings, CourierCostMapping)
		}
		courier := &pgsql.Courier{
			ID:                 i,
			Name:               randomdata.SillyName(),
			CourierCostMapping: courierCostMappings,
		}
		buyer := &pgsql.Buyer{
			ID:              i,
			Name:            randomdata.SillyName(),
			Province:        "Jawa Barat",
			City:            "Bekasi",
			Address:         "Jl Candrabaga " + index,
			Postcode:        randomdata.PostalCode("SE"),
			Latitude:        "",
			Longitude:       "",
			ProfileImageURL: "https://pbs.twimg.com/profile_images/1407698877914902530/Uy5uB6Qb_400x400.jpg",
		}
		db.Create(seller)
		db.Create(courier)
		db.Create(buyer)
	}
}

func main() {
	db := connectDB()

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "migrate":
			migration(db)
			break
		case "seed":
			seed(db)
			break
		}
		return
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
