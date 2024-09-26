package main

import (
	"database/sql"
	"log"

	"github.com/desafio-ll/cmd/server/handler"
	"github.com/desafio-ll/internal/dentist"
	"github.com/desafio-ll/internal/patient"
	"github.com/desafio-ll/pkg/store"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al intentar cargar archivo .env")
	}

	db, err := sql.Open("mysql", "root:root@/my_db")
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	storage := store.NewSqlStore(db)
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

	repoDentist := dentist.NewRepository(storage)
	serviceDentist := dentist.NewService(repoDentist)
	dentistHandler := handler.NewDentistHandler(serviceDentist)
	dentist := r.Group("/dentist")
	{
		dentist.GET(":id", dentistHandler.GetDentistByID())
		dentist.POST("", dentistHandler.Post())
		dentist.PUT(":id", dentistHandler.Put())
		dentist.PATCH(":id", dentistHandler.Patch())
		dentist.DELETE(":id", dentistHandler.Delete())
	}

	repoPatient := patient.NewRepository(storage)
	servicePatient := patient.NewService(repoPatient)
	patientHandler := handler.NewPatientHandler(servicePatient)
	patient := r.Group("/patient")
	{
		patient.GET(":id", patientHandler.GetPatientByID())
		patient.POST("", patientHandler.Post())
		patient.PUT(":id", patientHandler.Put())
		patient.PATCH(":id", patientHandler.Patch())
		patient.DELETE(":id", patientHandler.Delete())
	}
	r.Run(":8080")

}
