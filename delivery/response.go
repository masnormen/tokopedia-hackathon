package delivery

type SearchResponse struct {
	Success bool                     `json:"success"`
	Data    []map[string]interface{} `json:"data"`
}
