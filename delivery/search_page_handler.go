package delivery

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/masnormen/tokopedia-hackathon/usecase"
	"net/http"
)

type SearchPageHandler interface {
	Search()
	Sort()
}

type searchPageHandler struct {
	searchPageUsecase usecase.SearchPageUsecase
}

func (h *searchPageHandler) Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resp := SearchResponse{
		Success: true,
		Data: []SearchData{
			{
				ProductName:     "sabun dettol",
				Price:           12000,
				Rating:          4.5,
				TotalSales:      122,
				ProductImageURL: "google.com",
				SellerName:      "Yudit",
				SellerBadge:     "PRO",
				SellerCity:      "Bekasi",
				ShippingCourier: "JNE",
				ShippingCost:    10000,
			},
		},
	}

	json.NewEncoder(w).Encode(resp)
}

func (h *searchPageHandler) Sort(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resp := SearchResponse{
		Success: true,
		Data: []SearchData{
			{
				ProductName:     "sabun dettol",
				Price:           12000,
				Rating:          4.5,
				TotalSales:      122,
				ProductImageURL: "google.com",
				SellerName:      "Yudit",
				SellerBadge:     "PRO",
				SellerCity:      "Bekasi",
				ShippingCourier: "JNE",
				ShippingCost:    10000,
			},
		},
	}

	json.NewEncoder(w).Encode(resp)

}

func NewSearchPageHandler(r *mux.Router, searchPageUsecase usecase.SearchPageUsecase) {
	handler := searchPageHandler{searchPageUsecase: searchPageUsecase}
	r.HandleFunc("/api/v1/search", handler.Search).Methods("GET")
	r.HandleFunc("/api/v1/sort", handler.Sort).Methods("GET")
}
