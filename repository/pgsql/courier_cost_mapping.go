package pgsql

type CourierCostMapping struct {
	Id                      int     `pg:",pk" json:"id"`
	CourierId               int     `json:"courier_id"`
	ServiceName             string  `json:"service_name"`
	PostcodeCitySource      string  `json:"postcode_city_source"`
	PostcodeCityDestination string  `json:"postcode_city_destination"`
	BasePrice               int     `json:"base_price"`
	CoefWeight              float64 `json:"coef_weight"`
}
