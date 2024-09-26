package domain

type Patient struct {
	Id        int    `json:"id"`
	Nombre    string `json:"nombre" binding:"required"`
	Apellido  string `json:"apellido" binding:"required"`
	Domicilio string `json:"domicilio" binding:"required"`
	Dni       int    `json:"dni" binding:"required"`
	FechaAlta string `json:"fecha_alta" binding:"required"`
}
