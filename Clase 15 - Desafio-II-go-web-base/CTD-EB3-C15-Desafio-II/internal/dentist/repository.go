package dentist

import (
	"errors"

	"github.com/desafio-ll/internal/domain"
	"github.com/desafio-ll/pkg/store"
)

type RepositoryDentist interface {
	// GetByID busca un dentista por su id
	GetDentistByID(id int) (domain.Dentist, error)
	// Create agrega un nuevo dentista
	CreateDentist(p domain.Dentist) (domain.Dentist, error)
	// Update actualiza un dentista
	UpdateDentist(id int, p domain.Dentist) (domain.Dentist, error)
	// Delete elimina un dentista
	DeleteDentist(id int) error
}

type repositoryDentist struct {
	storage store.StoreInterface
}

// NewRepository crea un nuevo repositorio
func NewRepository(storage store.StoreInterface) RepositoryDentist {
	return &repositoryDentist{storage}
}

func (r *repositoryDentist) GetDentistByID(id int) (domain.Dentist, error) {
	dentist, err := r.storage.GetDentistByID(id)
	if err != nil {
		return domain.Dentist{}, errors.New("dentist not found")
	}
	return dentist, nil

}

func (r *repositoryDentist) CreateDentist(d domain.Dentist) (domain.Dentist, error) {
	if r.storage.ExistsDentist(d.Matricula) {
		return domain.Dentist{}, errors.New("there is already a dentist with this registration")
	}
	err := r.storage.CreateDentist(d)
	if err != nil {
		return domain.Dentist{}, errors.New("error creating dentist")
	}
	return d, nil
}

func (r *repositoryDentist) UpdateDentist(id int, d domain.Dentist) (domain.Dentist, error) {
	err := r.storage.UpdateDentist(d)
	if err != nil {
		return domain.Dentist{}, errors.New("error updating dentist")
	}
	return d, nil
}

func (r *repositoryDentist) DeleteDentist(id int) error {
	err := r.storage.DeleteDentist(id)
	if err != nil {
		return err
	}
	return nil
}
