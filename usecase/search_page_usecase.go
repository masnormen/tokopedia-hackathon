package usecase

import (
	"fmt"
	"github.com/masnormen/tokopedia-hackathon/repository/mysql"
)

type SearchPageUsecase interface {
	Search(searchQuery string, postCode string, sort string) ([]map[string]interface{}, error)
	Sort()
}

type searchPageUsecase struct {
	productOrm mysql.ProductOrm
}

func (s *searchPageUsecase) Search(searchQuery string, postCode string, sort string) ([]map[string]interface{}, error) {
	res, err := s.productOrm.Search(searchQuery, postCode, sort)
	if err != nil {
		fmt.Errorf("[SearchPageUsecase][Search] Error")
		return nil, err
	}

	return res, nil
}

func (s *searchPageUsecase) Sort() {
	panic("implement me")
}

func NewSearchPageUsecase(productOrm mysql.ProductOrm) SearchPageUsecase {
	return &searchPageUsecase{productOrm: productOrm}
}
