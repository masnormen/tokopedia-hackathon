package pgsql

type Courier struct {
	Id                 int                   `pg:",pk" json:"id"`
	Name               string                `json:"name"`
	CourierCostMapping []*CourierCostMapping `pg:"rel:has-many" json:"courier_cost_mapping"`
}
