package product

import (
	"gitlab.com/tomas.pereyra/product-service/internal/domain"
)

type Repository struct {
	database map[string]domain.Product
}

func NewRepository() *Repository {
	var productA = domain.Product{"1", "Teclado", 5000, "www.mys3bucket.com/1"}
	var productB = domain.Product{"1", "Notebook", 200000, "www.mys3bucket.com/2"}
	var productsMap = map[string]domain.Product{
		"1": productA,
		"2": productB,
	}

	return &Repository{database: productsMap}

}
func (repository *Repository) FindById(id string) domain.Product {

	return repository.database[id]
}

func (repository *Repository) Save(product domain.Product) domain.Product {
	repository.database[product.Id] = product
	return product
}
