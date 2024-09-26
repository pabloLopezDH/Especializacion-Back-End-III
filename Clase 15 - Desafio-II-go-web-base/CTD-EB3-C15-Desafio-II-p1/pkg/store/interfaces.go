package store

import "github.com/desafio-ll/internal/domain"

type StoreInterface interface {

	// GetDentistByID devuelve un dentista por su id
	GetDentistByID(id int) (domain.Dentist, error)
	// CreateDentist agrega un nuevo dentista
	CreateDentist(dentist domain.Dentist) error
	// UpdateDentist actualiza un dentista
	UpdateDentist(dentist domain.Dentist) error
	// DeleteDentist elimina un dentista
	DeleteDentist(id int) error
	// ExistsDentist verifica si un dentista existe
	ExistsDentist(matricula string) bool

	// GetPatientByID devuelve un paciente por su id
	GetPatientByID(id int) (domain.Patient, error)
	// CreatePatient agrega un nuevo paciente
	CreatePatient(patient domain.Patient) error
	// UpdatePatient actualiza un paciente
	UpdatePatient(patient domain.Patient) error
	// DeletePatient elimina un paciente
	DeletePatient(id int) error
	// ExistsPatient verifica si un paciente existe
	ExistsPatient(dni int) bool
}
