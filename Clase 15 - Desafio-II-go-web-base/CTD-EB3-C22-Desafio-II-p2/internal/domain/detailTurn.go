package domain

type DetailTurn struct {
	FechaHora         string `json:"fecha_hora" binding:"required"`
	Descripcion       string `json:"descripcion" binding:"required"`
	NombrePaciente    string `json:"nombre_paciente" binding:"required"`
	ApellidoPaciente  string `json:"apellido_paciente" binding:"required"`
	DniPaciente       int    `json:"dni_paciente" binding:"required"`
	FechaAltaPaciente string `json:"fecha_alta_paciente" binding:"required"`
	NombreDentista    string `json:"nombre_dentista" binding:"required"`
	ApellidoDentista  string `json:"apellido_dentista" binding:"required"`
	Matricula         int    `json:"matricula" binding:"required"`
}
