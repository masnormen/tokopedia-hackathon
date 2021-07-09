package delivery

import (
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
	panic("implement me")
}

func (h *searchPageHandler) Sort(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func NewSearchPageHandler(r *mux.Router, searchPageUsecase usecase.SearchPageUsecase) {
	handler := searchPageHandler{searchPageUsecase: searchPageUsecase}
	r.HandleFunc("/api/v1/search", handler.Search).Methods("GET")
	r.HandleFunc("/api/v1/sort", handler.Sort).Methods("GET")
}
