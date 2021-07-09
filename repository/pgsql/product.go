package pgsql

import (
	"gorm.io/gorm"
)

type Product struct {
	ID              int     `json:"id"`
	Name            string  `gorm:"index" json:"name"`
	Price           int     `json:"price"`
	Rating          float64 `json:"rating"`
	TotalSales      int     `json:"total_sales"`
	Weight          float64 `json:"weight"`
	ProductImageURL string  `json:"product_image_url"`
	SellerID        int     `json:"seller_id"`
	Seller          *Seller `json:"seller"`
	IsBebasOngkir   bool    `json:"is_bebas_ongkir"`
}

type ProductOrm interface {
	Search(query string, postCode string, sort string) ([]map[string]interface{}, error)
}

type productOrm struct {
	db *gorm.DB
}

func (p *productOrm) Search(searchQuery string, postCode string, sort string) ([]map[string]interface{}, error) {
	var result []map[string]interface{}

	query := `SELECT p.id AS id, p.name AS ProductName, p.price AS Price, TRUNCATE(p.rating, 1) AS Rating, p.total_sales AS TotalSales, p.product_image_url AS ProductImageURL, s.name AS SellerName, s.city AS SellerCity, 
			IF (p.is_bebas_ongkir, TRUNCATE(MIN(p.weight * ssc.coef_weight + ssc.base_price - 10000), 0), TRUNCATE(MIN(p.weight * ssc.coef_weight + ssc.base_price), 0))  AS ShippingCost
			FROM products AS p
			INNER JOIN sellers AS s ON p.seller_id = s.id
			INNER JOIN courier_cost_mappings AS ssc ON s.postcode = ssc.postcode_city_source AND ssc.postcode_city_destination = ?
			WHERE p.name LIKE CONCAT('%', ?, '%')
			GROUP BY p.id
			`

	switch sort {
	case "lowestprice":
		query = query + `ORDER BY p.price ASC;`
	case "lowesttco":
		query = query + `ORDER BY (ShippingCost + p.price) ASC;`
	default:
		query = query + `ORDER BY rating DESC, (ShippingCost + p.price) ASC, total_sales DESC;`
	}

	err := p.db.Raw(query, postCode, searchQuery).Scan(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil

}

func NewProductOrm(db *gorm.DB) ProductOrm {
	return &productOrm{db: db}
}
