package domain

type Turn struct {
	Id          int    `json:"id"`
	PacienteID  int    `json:"paciente_id" binding:"required"`
	DentistaID  int    `json:"dentista_id" binding:"required"`
	FechaHora   string `json:"fecha_hora" binding:"required"`
	Descripcion string `json:"descripcion" binding:"required"`
}
