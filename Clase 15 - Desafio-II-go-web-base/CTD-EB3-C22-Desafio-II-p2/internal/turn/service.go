package turn

import (
	"github.com/desafio-ll/internal/domain"
)

type Service interface {
	// GetTurnByID devuelve un turno por su id
	GetTurnByID(id int) (domain.Turn, error)
	// CreateTurn agrega un nuevo turno
	CreateTurn(p domain.Turn) (domain.Turn, error)
	// CreateBodyTurn turno por ID de paciente y ID de dentista
	CreateBodyTurn(p domain.BodyTurn) (domain.BodyTurn, error)
	// UpdateTurn actualiza un turno
	UpdateTurn(id int, p domain.Turn) (domain.Turn, error)
	// DeleteTurn elimina un turno
	DeleteTurn(id int) error
	// GetTurnByDniPatient devuelve un turno por dni paciente
	GetTurnByDniPatient(dni int) (domain.DetailTurn, error)
}

type service struct {
	r RepositoryTurn
}

// NewService crea un nuevo servicio
func NewService(r RepositoryTurn) Service {
	return &service{r}
}

func (s *service) GetTurnByID(id int) (domain.Turn, error) {
	p, err := s.r.GetTurnByID(id)
	if err != nil {
		return domain.Turn{}, err
	}
	return p, nil
}

func (s *service) CreateTurn(d domain.Turn) (domain.Turn, error) {
	d, err := s.r.CreateTurn(d)
	if err != nil {
		return domain.Turn{}, err
	}
	return d, nil
}

func (s *service) CreateBodyTurn(d domain.BodyTurn) (domain.BodyTurn, error) {
	d, err := s.r.CreateBodyTurn(d)
	if err != nil {
		return domain.BodyTurn{}, err
	}
	return d, nil
}

func (s *service) UpdateTurn(id int, u domain.Turn) (domain.Turn, error) {
	p, err := s.r.GetTurnByID(id)
	if err != nil {
		return domain.Turn{}, err
	}
	if u.PacienteID != 0 {
		p.PacienteID = u.PacienteID
	}
	if u.DentistaID != 0 {
		p.DentistaID = u.DentistaID
	}
	if u.FechaHora != "" {
		p.FechaHora = u.FechaHora
	}
	if u.Descripcion != "" {
		p.Descripcion = u.Descripcion
	}
	p, err = s.r.UpdateTurn(id, p)
	if err != nil {
		return domain.Turn{}, err
	}
	return p, nil
}

func (s *service) DeleteTurn(id int) error {
	err := s.r.DeleteTurn(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) GetTurnByDniPatient(dni int) (domain.DetailTurn, error) {
	p, err := s.r.GetTurnByDniPatient(dni)
	if err != nil {
		return domain.DetailTurn{}, err
	}
	return p, nil
}
