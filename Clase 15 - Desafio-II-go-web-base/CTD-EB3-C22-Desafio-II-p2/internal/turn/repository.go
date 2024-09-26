package turn

import (
	"errors"

	"github.com/desafio-ll/internal/domain"
	"github.com/desafio-ll/pkg/store"
)

type RepositoryTurn interface {
	// GetTurnByID busca un turno por su id
	GetTurnByID(id int) (domain.Turn, error)
	// CreateTurn agrega un nuevo turno
	CreateTurn(p domain.Turn) (domain.Turn, error)
	// UpdateTurn actualiza un turno
	UpdateTurn(id int, p domain.Turn) (domain.Turn, error)
	// DeleteTurn elimina un turno
	DeleteTurn(id int) error
	// CreateBodyTurn agrega un turno por ID de paciente y ID de dentista
	CreateBodyTurn(p domain.BodyTurn) (domain.BodyTurn, error)
	// GetTurnByDniPatient devuelve un turno por dni paciente
	GetTurnByDniPatient(dni int) (domain.DetailTurn, error)
}

type repositoryTurn struct {
	storage store.StoreInterface
}

// NewRepository crea un nuevo repositorio
func NewRepository(storage store.StoreInterface) RepositoryTurn {
	return &repositoryTurn{storage}
}

func (r *repositoryTurn) GetTurnByID(id int) (domain.Turn, error) {
	turn, err := r.storage.GetTurnByID(id)
	if err != nil {
		return domain.Turn{}, errors.New("turn not found")
	}
	return turn, nil
}

func (r *repositoryTurn) CreateTurn(d domain.Turn) (domain.Turn, error) {
	err := r.storage.CreateTurn(d)
	if err != nil {
		return domain.Turn{}, errors.New("error creating turn")
	}
	return d, nil
}

func (r *repositoryTurn) CreateBodyTurn(d domain.BodyTurn) (domain.BodyTurn, error) {
	err := r.storage.CreateBodyTurn(d)
	if err != nil {
		return domain.BodyTurn{}, errors.New("error creating turn")
	}
	return d, nil
}

func (r *repositoryTurn) UpdateTurn(id int, d domain.Turn) (domain.Turn, error) {
	err := r.storage.UpdateTurn(d)
	if err != nil {
		return domain.Turn{}, errors.New("error updating turn")
	}
	return d, nil
}

func (r *repositoryTurn) DeleteTurn(id int) error {
	err := r.storage.DeleteTurn(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repositoryTurn) GetTurnByDniPatient(dni int) (domain.DetailTurn, error) {
	turn, err := r.storage.GetTurnByDniPatient(dni)
	if err != nil {
		return domain.DetailTurn{}, errors.New("turn not found")
	}
	return turn, nil

}
