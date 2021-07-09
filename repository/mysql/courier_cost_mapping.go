package mysql

type CourierCostMapping struct {
	ID                      int      `json:"id"`
	CourierID               int      `json:"courier_id"`
	Courier                 *Courier `json:"courier"`
	ServiceName             string   `json:"service_name"`
	PostcodeCitySource      string   `gorm:"index:idx_member" json:"postcode_city_source"`
	PostcodeCityDestination string   `gorm:"index:idx_member" json:"postcode_city_destination"`
	BasePrice               int      `json:"base_price"`
	CoefWeight              int      `json:"coef_weight"`
}
