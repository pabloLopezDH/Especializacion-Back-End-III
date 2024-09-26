package handler

import (
	"errors"
	"os"
	"strconv"

	"github.com/desafio-ll/internal/domain"
	"github.com/desafio-ll/internal/turn"
	"github.com/desafio-ll/pkg/web"
	"github.com/gin-gonic/gin"
)

type turnHandler struct {
	s turn.Service
}

// NewTurnHandler crea un nuevo controller para turnos
func NewTurnHandler(s turn.Service) *turnHandler {
	return &turnHandler{
		s: s,
	}
}

// GetTurnByID obtiene un turno por id
func (h *turnHandler) GetTurnByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		turn, err := h.s.GetTurnByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("turn not found"))
			return
		}
		web.Success(c, 200, turn)
	}
}

// validateTurnEmptys valida que los campos no esten vacios
func validateTurnEmptys(turn *domain.Turn) (bool, error) {
	switch {
	case turn.PacienteID == 0 || turn.DentistaID == 0 || turn.FechaHora == "" || turn.Descripcion == "":
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}

// Post crea un nuevo turno por ID de paciente y ID de dentista
func (h *turnHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var turn domain.Turn
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		err := c.ShouldBindJSON(&turn)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateTurnEmptys(&turn)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.CreateTurn(turn)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p)
	}
}

// Post crea un nuevo turno por DNI paciente y MATRICULA de dentista
func (h *turnHandler) PostCreateTurn() gin.HandlerFunc {
	return func(c *gin.Context) {
		var bodyturn domain.BodyTurn
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		err := c.ShouldBindJSON(&bodyturn)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		p, err := h.s.CreateBodyTurn(bodyturn)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p)
	}
}

// Delete elimina un turno
func (h *turnHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		err = h.s.DeleteTurn(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 204, nil)
	}
}

// Put actualiza un turno
func (h *turnHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.s.GetTurnByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("turn not found"))
			return
		}
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		var turn domain.Turn
		err = c.ShouldBindJSON(&turn)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateTurnEmptys(&turn)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.UpdateTurn(id, turn)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}

// Patch actualiza un turno o alguno de sus campos
func (h *turnHandler) Patch() gin.HandlerFunc {
	type Request struct {
		PacienteID  int    `json:"paciente_id,omitempty"`
		DentistaID  int    `json:"dentista_id,omitempty"`
		FechaHora   string `json:"fecha_hora,omitempty"`
		Descripcion string `json:"descripcion,omitempty"`
	}
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		var r Request
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.s.GetTurnByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("turn not found"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		update := domain.Turn{
			PacienteID:  r.PacienteID,
			DentistaID:  r.DentistaID,
			FechaHora:   r.FechaHora,
			Descripcion: r.Descripcion,
		}
		p, err := h.s.UpdateTurn(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}

// GetTurnByDniPatient obtiene un turno por DNI paciente
func (h *turnHandler) GetTurnByDniPatient() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("dni")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid dni"))
			return
		}
		turn, err := h.s.GetTurnByDniPatient(id)
		if err != nil {
			web.Failure(c, 404, errors.New("turn not found"))
			return
		}
		web.Success(c, 200, turn)
	}
}
