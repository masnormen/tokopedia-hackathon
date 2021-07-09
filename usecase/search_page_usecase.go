package usecase

import "github.com/masnormen/tokopedia-hackathon/repository/pgsql"

type SearchPageUsecase interface {
	Search()
	Sort()
}

type searchPageUsecase struct {
	productOrm pgsql.ProductOrm
}

func (s *searchPageUsecase) Search() {
	panic("implement me")
}

func (s *searchPageUsecase) Sort() {
	panic("implement me")
}

func NewSearchPageUsecase(productOrm pgsql.ProductOrm) SearchPageUsecase {
	return &searchPageUsecase{productOrm: productOrm}
}
