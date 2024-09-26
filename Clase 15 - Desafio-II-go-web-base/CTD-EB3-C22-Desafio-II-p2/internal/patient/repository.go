package patient

import (
	"errors"

	"github.com/desafio-ll/internal/domain"
	"github.com/desafio-ll/pkg/store"
)

type RepositoryPatient interface {
	// GetPatientByID busca un paciente por su id
	GetPatientByID(id int) (domain.Patient, error)
	// CreatePatient agrega un nuevo paciente
	CreatePatient(p domain.Patient) (domain.Patient, error)
	// UpdatePatient actualiza un paciente
	UpdatePatient(id int, p domain.Patient) (domain.Patient, error)
	// DeletePatient elimina un paciente
	DeletePatient(id int) error
}

type repositoryPatient struct {
	storage store.StoreInterface
}

// NewRepository crea un nuevo repositorio
func NewRepository(storage store.StoreInterface) RepositoryPatient {
	return &repositoryPatient{storage}
}

func (r *repositoryPatient) GetPatientByID(id int) (domain.Patient, error) {
	patient, err := r.storage.GetPatientByID(id)
	if err != nil {
		return domain.Patient{}, errors.New("patient not found")
	}
	return patient, nil

}

func (r *repositoryPatient) CreatePatient(d domain.Patient) (domain.Patient, error) {
	if r.storage.ExistsPatient(d.Dni) {
		return domain.Patient{}, errors.New("there is already a patient with this dni")
	}
	err := r.storage.CreatePatient(d)
	if err != nil {
		return domain.Patient{}, errors.New("error creating patient")
	}
	return d, nil
}

func (r *repositoryPatient) UpdatePatient(id int, d domain.Patient) (domain.Patient, error) {
	err := r.storage.UpdatePatient(d)
	if err != nil {
		return domain.Patient{}, errors.New("error updating patient")
	}
	return d, nil
}

func (r *repositoryPatient) DeletePatient(id int) error {
	err := r.storage.DeletePatient(id)
	if err != nil {
		return err
	}
	return nil
}
