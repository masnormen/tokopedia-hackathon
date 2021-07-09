package pgsql

import (
	"gorm.io/gorm"
)

type Product struct {
	ID              int     `json:"id"`
	Name            string  `json:"name"`
	Price           int     `json:"price"`
	Rating          float64 `json:"rating"`
	TotalSales      int     `json:"total_sales"`
	ProductImageURL string  `json:"product_image_url"`
	SellerID        int     `json:"seller_id"`
	Seller          *Seller `json:"seller"`
}

type ProductOrm interface {
	Search()
}

type productOrm struct {
	db *gorm.DB
}

func (p *productOrm) Search() {
	panic("implement me")
}

func NewProductOrm(db *gorm.DB) ProductOrm {
	return &productOrm{db: db}
}
