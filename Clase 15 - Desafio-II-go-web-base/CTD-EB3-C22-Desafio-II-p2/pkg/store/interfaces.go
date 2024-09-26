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
	ExistsDentist(matricula int) bool

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

	// GetTurnByID devuelve un turno por su id
	GetTurnByID(id int) (domain.Turn, error)
	// CreateTurn agrega un nuevo turno
	CreateTurn(turn domain.Turn) error
	// UpdateTurn actualiza un turno
	UpdateTurn(turn domain.Turn) error
	// DeleteTurn elimina un turno
	DeleteTurn(id int) error

	// GetPatientByDni devuelve un id de un paciente por DNI
	GetPatientByDni(dni int) (int, error)
	// GetDentistByMatricula devuelve un id de un dentista por MATRICULA
	GetDentistByMatricula(matricula int) (int, error)

	// CreateBodyTurn agrega un turno por ID de paciente y ID de dentista
	CreateBodyTurn(turn domain.BodyTurn) error
	// GetTurnByDniPatient agrega nuevo turno por DNI paciente y MATRICULA de dentista
	GetTurnByDniPatient(dni int) (domain.DetailTurn, error)
}
