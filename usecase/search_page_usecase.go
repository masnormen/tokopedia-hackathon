package usecase

import (
	"fmt"
	"github.com/masnormen/tokopedia-hackathon/repository/pgsql"
)

type SearchPageUsecase interface {
	Search(searchQuery string, postCode string) ([]map[string]interface{}, error)
	Sort()
}

type searchPageUsecase struct {
	productOrm pgsql.ProductOrm
}

func (s *searchPageUsecase) Search(searchQuery string, postCode string) ([]map[string]interface{}, error) {
	res, err := s.productOrm.Search(searchQuery, postCode)
	if err != nil {
		fmt.Errorf("[SearchPageUsecase][Search] Error")
		return nil, err
	}

	return res, nil
}

func (s *searchPageUsecase) Sort() {
	panic("implement me")
}

func NewSearchPageUsecase(productOrm pgsql.ProductOrm) SearchPageUsecase {
	return &searchPageUsecase{productOrm: productOrm}
}
