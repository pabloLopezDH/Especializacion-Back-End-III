package patient

import (
	"github.com/desafio-ll/internal/domain"
)

type Service interface {
	// GetPatientByID busca un paciente por su id
	GetPatientByID(id int) (domain.Patient, error)
	// CreatePatient agrega un nuevo paciente
	CreatePatient(p domain.Patient) (domain.Patient, error)
	// UpdatePatient actualiza un paciente
	UpdatePatient(id int, p domain.Patient) (domain.Patient, error)
	// DeletePatient elimina un paciente
	DeletePatient(id int) error
}

type service struct {
	r RepositoryPatient
}

// NewService crea un nuevo servicio
func NewService(r RepositoryPatient) Service {
	return &service{r}
}

func (s *service) GetPatientByID(id int) (domain.Patient, error) {
	p, err := s.r.GetPatientByID(id)
	if err != nil {
		return domain.Patient{}, err
	}
	return p, nil
}

func (s *service) CreatePatient(d domain.Patient) (domain.Patient, error) {
	d, err := s.r.CreatePatient(d)
	if err != nil {
		return domain.Patient{}, err
	}
	return d, nil
}
func (s *service) UpdatePatient(id int, u domain.Patient) (domain.Patient, error) {
	p, err := s.r.GetPatientByID(id)
	if err != nil {
		return domain.Patient{}, err
	}
	if u.Nombre != "" {
		p.Nombre = u.Nombre
	}
	if u.Apellido != "" {
		p.Apellido = u.Apellido
	}
	if u.Domicilio != "" {
		p.Domicilio = u.Domicilio
	}
	if u.Dni != 0 {
		p.Dni = u.Dni
	}
	if u.FechaAlta != "" {
		p.FechaAlta = u.FechaAlta
	}
	p, err = s.r.UpdatePatient(id, p)
	if err != nil {
		return domain.Patient{}, err
	}
	return p, nil
}

func (s *service) DeletePatient(id int) error {
	err := s.r.DeletePatient(id)
	if err != nil {
		return err
	}
	return nil
}
