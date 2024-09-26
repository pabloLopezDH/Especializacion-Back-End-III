package handler

import (
	"errors"
	"os"
	"strconv"

	"github.com/desafio-ll/internal/domain"
	"github.com/desafio-ll/internal/patient"
	"github.com/desafio-ll/pkg/web"
	"github.com/gin-gonic/gin"
)

type patientHandler struct {
	s patient.Service
}

// NewPatientHandler crea un nuevo controller para paciente
func NewPatientHandler(s patient.Service) *patientHandler {
	return &patientHandler{
		s: s,
	}
}

// GetPatientByID obtiene un paciente por id
func (h *patientHandler) GetPatientByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		patient, err := h.s.GetPatientByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("patient not found"))
			return
		}
		web.Success(c, 200, patient)
	}
}

// validatePatientEmptys valida que los campos no esten vacios
func validatePatientEmptys(patient *domain.Patient) (bool, error) {
	switch {
	case patient.Nombre == "" || patient.Apellido == "" || patient.Domicilio == "" || patient.Dni == 0 || patient.FechaAlta == "":
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}

// Post crea un nuevo paciente
func (h *patientHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var patient domain.Patient
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		err := c.ShouldBindJSON(&patient)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validatePatientEmptys(&patient)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.CreatePatient(patient)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p)
	}
}

// Delete elimina un paciente
func (h *patientHandler) Delete() gin.HandlerFunc {
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
		err = h.s.DeletePatient(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 204, nil)
	}
}

// Put actualiza un paciente
func (h *patientHandler) Put() gin.HandlerFunc {
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
		_, err = h.s.GetPatientByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("patient not found"))
			return
		}
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		var patient domain.Patient
		err = c.ShouldBindJSON(&patient)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validatePatientEmptys(&patient)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.UpdatePatient(id, patient)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}

// Patch actualiza un paciente o alguno de sus campos
func (h *patientHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Nombre    string `json:"nombre,omitempty"`
		Apellido  string `json:"apellido,omitempty"`
		Domicilio string `json:"domicilio,omitempty"`
		Dni       int    `json:"dni,omitempty"`
		FechaAlta string `json:"fecha_alta,omitempty"`
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
		_, err = h.s.GetPatientByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("patient not found"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		update := domain.Patient{
			Nombre:    r.Nombre,
			Apellido:  r.Apellido,
			Domicilio: r.Domicilio,
			Dni:       r.Dni,
			FechaAlta: r.FechaAlta,
		}
		p, err := h.s.UpdatePatient(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}
