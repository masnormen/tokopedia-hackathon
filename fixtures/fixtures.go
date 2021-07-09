package fixtures

import (
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"github.com/masnormen/tokopedia-hackathon/repository/mysql"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
)

var arrayOfLink = []map[string]interface{}{
	{
		"images":   "https://images.tokopedia.net/img/cache/200-square/product-1/2020/1/31/619071290/619071290_4cade089-de33-4add-95c1-6895e557a4c4_896_896.jpg.webp?ect=4g",
		"name":     "[PREMIUM] SEPATU VANS OLD SCKOOL DT",
		"city":     "Depok",
		"postcode": "17601",
	},
	{
		"images":   "https://images.tokopedia.net/img/cache/200-square/VqbcmM/2021/1/27/62e7f29d-2b7b-413d-92bf-c70cd9b0b831.jpg.webp?ect=4g",
		"name":     "Terbaru!!! Sepatu Pria Fashione Sport AIFF 720 Sepatu Pria Running Sep",
		"city":     "Bogor",
		"postcode": "17602",
	},
	{
		"images": "https://images.tokopedia.net/img/cache/200-square/VqbcmM/2021/6/23/642390c7-f340-4869-96bd-ab20458c8a22.jpg.webp?ect=4g",
		"name":   "Sepatu Onitsuka Tiger Mexico66 GRD-R Original Leather Guarantee JAPAN - Magenta, 36",
		"city":   "bandung",
	},
	{
		"images": "https://images.tokopedia.net/img/cache/200-square/VqbcmM/2021/1/21/3c6aeb3f-0413-4cb6-8983-deb12e4c4fcc.jpg.webp?ect=4g",
		"name":   "Sepatu Pria Nike Air Force 1 All White 39 44 OBRAL MURAH",
		"city":   "jakarta selatan",
	},
	{
		"images": "https://images.tokopedia.net/img/cache/200-square/product-1/2020/3/10/727158633/727158633_53c547c9-6b16-4ac4-9c49-a1452476e1d6_1024_1024.jpg.webp?ect=4g",
		"name":   "SEPATU JOHNSON GALAXY PRO HIGH BLACK WHITE ORIGINAL INDONESIA 38-43",
		"city":   "jakarta selatan",
	},
	{
		"images": "https://images.tokopedia.net/img/cache/200-square/VqbcmM/2021/4/12/c7ef317d-5809-46c4-af29-5d42addb9023.jpg.webp?ect=4g",
		"name":   "Larocking - Venta Vol 1 Hitam | Sepatu Sneakers Running Gym Shoes",
		"city":   "bekasi",
	},
	{
		"images": "https://images.tokopedia.net/img/cache/200-square/VqbcmM/2020/9/9/83209a61-8917-4abb-a6b3-9fb1b0d3ff44.jpg.webp?ect=4g",
		"name":   "SEPATU CASUAL CONVERS ALLSTAR 2 PREMIUM MURAH",
		"city":   "Mojokerto",
	},
	{
		"images": "https://images.tokopedia.net/img/cache/200-square/VqbcmM/2021/4/8/85ecd187-6af8-49c8-bec5-02294d6ce364.jpg.webp?ect=4g",
		"name":   "NING Sepatu Air/ Sepatu Renang/ Sepatu Pantai atau HitamAbuKuning ZB333 - 37",
		"city":   "medan",
	},
	{
		"images": "https://images.tokopedia.net/img/cache/200-square/VqbcmM/2021/4/3/e9fe214e-d493-40a6-baba-cdc8c33fbe19.jpg.webp?ect=4g",
		"name":   "Sepatu VENTELA ALL Is WELL (Gading Marten) EVIL Series",
		"city":   "jakarta timur",
	},
	{
		"images": "https://images.tokopedia.net/img/cache/200-square/VqbcmM/2020/10/27/6717f23c-21ac-4691-a6b1-fcfd315d09c9.jpg.webp?ect=4g",
		"name":   "Sepatu Ventela Basic Low Black Natural Original Made In Indonesia",
		"city":   "jakarta selatan",
	},
	{
		"images": "https://images.tokopedia.net/img/cache/200-square/VqbcmM/2021/2/3/c8e1637a-ca2c-43e9-a3fe-b5415c8b1088.jpg.webp?ect=4g",
		"name":   "Sepatu converse 70s high white parchment - Putih, 36",
		"city":   "jakarta selatan",
	},
	{
		"images": "https://images.tokopedia.net/img/cache/200-square/product-1/2017/5/20/1733991/1733991_db5e7d99-d2a2-4f3d-a982-9136f9ca9c52.jpg.webp?ect=4g",
		"name":   "Sepatu Vans Oldskool/Vans cowok/vans cewek/vans Premium original",
		"city":   "jakarta utara",
	},
	{
		"images": "https://images.tokopedia.net/img/cache/200-square/product-1/2020/8/4/387156715/387156715_181ba6ab-7527-449e-a7bf-0dab2cd1692d_687_687.jpg.webp?ect=4g",
		"name":   "Sepatu Nike Air Jordan 4 KAWS Retro Cool Grey Premium Original Sneaker",
		"city":   "jakarta timur",
	},
	{
		"images": "https://images.tokopedia.net/img/cache/200-square/VqbcmM/2020/9/2/acc51726-04a4-46c1-b798-90d704f4511c.jpg.webp?ect=4g",
		"name":   "Sepatu Patrobas Equip Low (ORIGINAL) 2020",
		"city":   "jakarta timur",
	},
	{
		"images": "https://images.tokopedia.net/img/cache/200-square/VqbcmM/2021/1/21/3c6aeb3f-0413-4cb6-8983-deb12e4c4fcc.jpg.webp?ect=4g",
		"name":   "Sepatu Pria Nike Air Force 1 All White 39 44 OBRAL MURAH",
		"city":   "jakarta selatan",
	},
	{
		"images": "https://images.tokopedia.net/img/cache/200-square/VqbcmM/2021/7/2/16b69ef3-b5a6-4dbf-a0ef-6c06e1f93ad9.jpg.webp?ect=4g",
		"name":   "SEPATU BRANDED SNEAKERS PRIA IMPORT PR4D4 MIRROR 1:1 ORI TERMURAH #26",
		"city":   "jakarta utara",
	},
	{
		"images": "https://images.tokopedia.net/img/cache/200-square/VqbcmM/2021/6/25/60ce1679-0124-40b2-b3e6-903bbc901736.png.webp?ect=4g",
		"name":   "PHM Shoes Sepatu Pria Sneakers Import Sepatu Olahraga Kasual PHM103",
		"city":   "jakarta barat",
	},
	{
		"images": "https://images.tokopedia.net/img/cache/200-square/VqbcmM/2021/2/17/143f5e5c-9763-4fea-9248-319f4475aa01.jpg.webp?ect=4g",
		"name":   "Sepatu Casual Pria Kuzatura NC KZS 225",
		"city":   "bandung",
	},
	{
		"images": "https://images.tokopedia.net/img/cache/200-square/VqbcmM/2021/5/31/fe739d63-93d6-432d-8646-0850bf1277d4.jpg.webp?ect=4g",
		"name":   "Sepatu Pria Boots High Casual Kulit Sintetis Joemen J85 - Hitam, 39",
		"city":   "jakarta barat",
	},
	{
		"images": "https://images.tokopedia.net/img/cache/200-square/product-1/2017/5/17/2170221/2170221_aaf63b3d-85d6-49bd-adcf-db750ab58742.jpg.webp?ect=4g",
		"name":   "Sepatu boot air karet Jeep 9001",
		"city":   "sidoarjo",
	},
	{
		"images": "https://images.tokopedia.net/img/cache/200-square/product-1/2020/1/9/69327184/69327184_39219ec7-6936-4464-ac2a-4aeaa25a60d1_1280_1280.webp?ect=4g",
		"name":   "sepatu pria crocodile armor boots safety tracking",
		"city":   "bandung",
	},
	{
		"images": "https://images.tokopedia.net/img/cache/200-square/VqbcmM/2020/11/20/50b580e0-89f9-4f59-a4e9-ea19c7cc389b.jpg.webp?ect=4g",
		"name":   "SEPATU PDL NINJA PRIA GEN 6 MERK BROTRHER",
		"city":   "bandung",
	},
	{
		"images": "https://images.tokopedia.net/img/cache/200-square/product-1/2019/11/25/256180439/256180439_ca3662bf-a211-4dae-9e7a-ca8dd3c63f5b_414_414.jpg.webp?ect=4g",
		"name":   "Sepatu Adidas Full Black Big size 45-47",
		"city":   "Mojokerto",
	},
}

var arrayOfPostcode = []string{"17601", "17602", "17603", "17604", "17605", "36131", "36132", "36133", "36134", "36135"}
var arrayOfCity = []string{"Kota Bekasi", "Kab. Bekasi", "Jakarta Selatan", "Jakarta Pusat", "Jakarta Utara", "Kota Baru", "Alam Barajo", "Jambi Timur", "Pasar Jambi", "Pelayangan"}

func Seed(db *gorm.DB) {
	for i := 1; i < 10; i++ {
		index := strconv.Itoa(i)
		products := make([]*mysql.Product, 0)
		for j := 1; j < 10; j++ {
			product := &mysql.Product{
				Name:            arrayOfLink[i+j]["name"].(string),
				Price:           rand.Intn(350000),
				Rating:          randomdata.Decimal(1, 5),
				TotalSales:      rand.Intn(1200),
				Weight:          randomdata.Decimal(1, 10),
				ProductImageURL: arrayOfLink[i+j]["images"].(string),
				SellerID:        i,
				IsBebasOngkir:   randomdata.Boolean(),
			}
			products = append(products, product)
		}
		seller := &mysql.Seller{
			ID:              i,
			Name:            randomdata.SillyName(),
			Province:        randomdata.ProvinceForCountry("62"),
			City:            arrayOfCity[i-1],
			Address:         "Jl Candrabaga " + index,
			Postcode:        arrayOfPostcode[i-1],
			Latitude:        "",
			Longitude:       "",
			ProfileImageURL: "https://pbs.twimg.com/profile_images/1407698877914902530/Uy5uB6Qb_400x400.jpg",
			Badge:           rand.Intn(3),
			Product:         products,
		}

		db.Create(seller)
	}

	buyer1 := &mysql.Buyer{
		ID:              1,
		Name:            randomdata.SillyName(),
		Province:        "Jawa Barat",
		City:            "Bekasi",
		Address:         "Jl Candrabaga",
		Postcode:        "17612",
		Latitude:        "",
		Longitude:       "",
		ProfileImageURL: "https://pbs.twimg.com/profile_images/1407698877914902530/Uy5uB6Qb_400x400.jpg",
	}
	db.Create(buyer1)

	buyer2 := &mysql.Buyer{
		ID:              2,
		Name:            randomdata.SillyName(),
		Province:        "Jambi",
		City:            "Kota Jambi",
		Address:         "Jl Candrabaga",
		Postcode:        "36138",
		Latitude:        "",
		Longitude:       "",
		ProfileImageURL: "https://pbs.twimg.com/profile_images/1407698877914902530/Uy5uB6Qb_400x400.jpg",
	}
	db.Create(buyer2)

	for i := 1; i < 10; i++ {
		courier := &mysql.Courier{
			ID:   i,
			Name: randomdata.SillyName(),
		}

		courierCostMappings := make([]*mysql.CourierCostMapping, 0)
		for j := 0; j < 5; j++ {
			courierCostMapping1 := &mysql.CourierCostMapping{
				CourierID:               i,
				ServiceName:             randomdata.SillyName(),
				PostcodeCitySource:      arrayOfPostcode[j],
				PostcodeCityDestination: "17612",
				BasePrice:               randomdata.Number(1000, 10000),
				CoefWeight:              randomdata.Number(1000, 10000),
			}
			courierCostMappings = append(courierCostMappings, courierCostMapping1)

			courierCostMapping2 := &mysql.CourierCostMapping{
				CourierID:               i,
				ServiceName:             randomdata.SillyName(),
				PostcodeCitySource:      arrayOfPostcode[j],
				PostcodeCityDestination: "36138",
				BasePrice:               randomdata.Number(30000, 100000),
				CoefWeight:              randomdata.Number(5000, 20000),
			}
			courierCostMappings = append(courierCostMappings, courierCostMapping2)
		}

		for j := 5; j < 10; j++ {
			courierCostMapping1 := &mysql.CourierCostMapping{
				CourierID:               i,
				ServiceName:             randomdata.SillyName(),
				PostcodeCitySource:      arrayOfPostcode[j],
				PostcodeCityDestination: "17612",
				BasePrice:               randomdata.Number(30000, 100000),
				CoefWeight:              randomdata.Number(5000, 20000),
			}
			courierCostMappings = append(courierCostMappings, courierCostMapping1)

			courierCostMapping2 := &mysql.CourierCostMapping{
				CourierID:               i,
				ServiceName:             randomdata.SillyName(),
				PostcodeCitySource:      arrayOfPostcode[j],
				PostcodeCityDestination: "36138",
				BasePrice:               randomdata.Number(1000, 10000),
				CoefWeight:              randomdata.Number(1000, 10000),
			}
			courierCostMappings = append(courierCostMappings, courierCostMapping2)

		}

		courier.CourierCostMapping = courierCostMappings
		db.Create(courier)
	}

}

func Migrate(db *gorm.DB) error {
	fmt.Println("Doing migration...")
	err := db.AutoMigrate(
		&mysql.Buyer{},
		&mysql.Courier{},
		&mysql.CourierCostMapping{},
		&mysql.Product{},
		&mysql.Seller{},
	)

	if err != nil {
		panic(err)
	}
	fmt.Println("Done")
	return nil
}
