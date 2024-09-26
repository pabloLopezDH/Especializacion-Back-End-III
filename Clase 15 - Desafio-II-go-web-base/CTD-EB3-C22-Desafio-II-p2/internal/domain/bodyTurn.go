package domain

type BodyTurn struct {
	Id                int    `json:"id"`
	DniPaciente       int    `json:"dni_paciente" binding:"required"`
	MatriculaDentista int    `json:"matricula_dentista" binding:"required"`
	FechaHora         string `json:"fecha_hora" binding:"required"`
	Descripcion       string `json:"descripcion" binding:"required"`
}
