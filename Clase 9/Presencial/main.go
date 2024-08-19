package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Persona struct {
	Nombre    string
	Apellido  string
	Edad      int
	Direccion string
	Telefono  string
	Activo    bool
}

func main() {
	//go mod init go-web
	//go get -u github.com/gin-gonic/gin

	router := gin.Default()

	var p Persona

	jsonData := `{"Nombre":"Juan", "Apellido":"Perez", "Edad": 25, "Direccion":"Calle falsa 123", "Telefono":"1155447788", "Activo": true}`

	if err := json.Unmarshal([]byte(jsonData), &p); err != nil {
		log.Fatal(err)
	}

	fmt.Println(p)

	router.GET("/persona", func(ctx *gin.Context) {
		panic("test")
		ctx.JSON(http.StatusOK, gin.H{
			"persona": p,
		})
	})

	router.GET("/persona1", func(ctx *gin.Context) {
		//panic("test")
		ctx.JSON(http.StatusOK, gin.H{
			"persona": p,
		})
	})

	router.Run(":8080")

}
