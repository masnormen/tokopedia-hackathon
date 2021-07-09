package delivery

import (
	"encoding/json"
	"fmt"
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
	v := r.URL.Query()

	searchQuery := v.Get("q")
	postCode := v.Get("from")
	sort := v.Get("sort")

	var resp SearchResponse
	w.Header().Set("Content-Type", "application/json")

	res, err := h.searchPageUsecase.Search(searchQuery, postCode, sort)
	if err != nil {
		fmt.Errorf("[SearchPageHandler][Search] Error")
		resp.Success = false
		resp.Data = nil
		json.NewEncoder(w).Encode(resp)
	}

	resp.Success = true
	resp.Data = res
	json.NewEncoder(w).Encode(resp)
}

func (h *searchPageHandler) Sort(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()

	searchQuery := v.Get("q")
	postCode := v.Get("from")
	sort := v.Get("sort")

	var resp SearchResponse
	w.Header().Set("Content-Type", "application/json")

	res, err := h.searchPageUsecase.Search(searchQuery, postCode, sort)
	if err != nil {
		fmt.Errorf("[SearchPageHandler][Search] Error")
		resp.Success = false
		resp.Data = nil
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp.Success = true
	resp.Data = res
	json.NewEncoder(w).Encode(resp)

}

func NewSearchPageHandler(r *mux.Router, searchPageUsecase usecase.SearchPageUsecase) {
	handler := searchPageHandler{searchPageUsecase: searchPageUsecase}
	r.HandleFunc("/api/v1/search", handler.Search).Methods("GET")
	r.HandleFunc("/api/v1/sort", handler.Sort).Methods("GET")
}
