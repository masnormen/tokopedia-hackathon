package pgsql

type CourierCostMapping struct {
	ID                      int      `json:"id"`
	CourierID               int      `json:"courier_id"`
	Courier                 *Courier `json:"courier"`
	ServiceName             string   `json:"service_name"`
	PostcodeCitySource      string   `json:"postcode_city_source"`
	PostcodeCityDestination string   `json:"postcode_city_destination"`
	BasePrice               int      `json:"base_price"`
	CoefWeight              float64  `json:"coef_weight"`
}
