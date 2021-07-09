package delivery

type SearchResponse struct {
	Success bool         `json:"success"`
	Data    []SearchData `json:"data"`
}

type SearchData struct {
	ProductName     string  `json:"ProductName"`
	Price           int     `json:"Price"`
	Rating          float64 `json:"Rating"`
	TotalSales      int     `json:"TotalSales"`
	ProductImageURL string  `json:"ProductImageURL"`
	SellerName      string  `json:"SellerName"`
	SellerBadge     string  `json:"SellerBadge"`
	SellerCity      string  `json:"SellerCity"`
	ShippingCourier string  `json:"ShippingCourier"`
	ShippingCost    int     `json:"ShippingCost"`
}
