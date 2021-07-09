package pgsql

type Seller struct {
	ID              int        `pg:",pk" json:"id"`
	Name            string     `json:"name"`
	Province        string     `json:"province"`
	City            string     `json:"city"`
	Address         string     `json:"address"`
	Postcode        string     `json:"postcode"`
	Latitude        string     `json:"latitude"`
	Longitude       string     `json:"longitude"`
	ProfileImageURL string     `json:"profile_image_url"`
	Badge           string     `json:"badge"`
	Product         []*Product `pg:"rel:has-many" json:"product"`
}
