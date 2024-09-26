package product

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/tomas.pereyra/product-service/internal/domain"
)

type Service struct {
	repository Repository
}

func NewService(r Repository) *Service {
	return &Service{r}
}
func (service *Service) FindProductById(ctxt *gin.Context) domain.Product {
	return service.repository.FindById(ctxt.Param("id"))
}

func (service *Service) SaveProduct(product domain.Product) domain.Product {
	return service.repository.Save(product)

}
