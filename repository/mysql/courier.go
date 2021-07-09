package mysql

type Courier struct {
	ID                 int                   `json:"id"`
	Name               string                `json:"name"`
	CourierCostMapping []*CourierCostMapping `json:"courier_cost_mapping"`
}
