package pgsql

import (
	"github.com/go-pg/pg/v10"
)

type Product struct {
	Id              int     `pg:",pk" json:"id"`
	Name            string  `json:"name"`
	Price           int     `json:"price"`
	Rating          float64 `json:"rating"`
	TotalSales      int     `json:"total_sales"`
	ProductImageURL string  `json:"product_image_url"`
	SellerId        int     `json:"seller_id"`
}

type ProductOrm interface {
	Search()
}

type productOrm struct {
	db *pg.DB
}

func (p *productOrm) Search() {
	panic("implement me")
}

func NewProductOrm(db *pg.DB) ProductOrm {
	return &productOrm{db: db}
}
