package dentist

import (
	"github.com/desafio-ll/internal/domain"
)

type Service interface {
	// GetDentistByID busca un dentista por su id
	GetDentistByID(id int) (domain.Dentist, error)
	// CreateDentist agrega un nuevo dentista
	CreateDentist(p domain.Dentist) (domain.Dentist, error)
	// UpdateDentist actualiza un dentista
	UpdateDentist(id int, p domain.Dentist) (domain.Dentist, error)
	// DeleteDentist elimina un dentista
	DeleteDentist(id int) error
}

type service struct {
	r RepositoryDentist
}

// NewService crea un nuevo servicio
func NewService(r RepositoryDentist) Service {
	return &service{r}
}

func (s *service) GetDentistByID(id int) (domain.Dentist, error) {
	p, err := s.r.GetDentistByID(id)
	if err != nil {
		return domain.Dentist{}, err
	}
	return p, nil
}

func (s *service) CreateDentist(d domain.Dentist) (domain.Dentist, error) {
	d, err := s.r.CreateDentist(d)
	if err != nil {
		return domain.Dentist{}, err
	}
	return d, nil
}
func (s *service) UpdateDentist(id int, u domain.Dentist) (domain.Dentist, error) {
	p, err := s.r.GetDentistByID(id)
	if err != nil {
		return domain.Dentist{}, err
	}
	if u.Nombre != "" {
		p.Nombre = u.Nombre
	}
	if u.Apellido != "" {
		p.Apellido = u.Apellido
	}
	if u.Matricula != 0 {
		p.Matricula = u.Matricula
	}
	p, err = s.r.UpdateDentist(id, p)
	if err != nil {
		return domain.Dentist{}, err
	}
	return p, nil
}

func (s *service) DeleteDentist(id int) error {
	err := s.r.DeleteDentist(id)
	if err != nil {
		return err
	}
	return nil
}
