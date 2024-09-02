package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

/*
● Verbo utilizado: GET, POST, PUT, etc.
● Fecha y hora: podemos  utilizar el paquete time.
● URL de la consulta: localhost:8080/products
● Tamaño en bytes: tamaño  de la consulta.
*/

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verb := ctx.Request.Method
		time1 := time.Now()
		url := ctx.Request.URL

		//Procesa el siguiente middleware
		ctx.Next()
		var size int
		if ctx.Writer != nil {
			size = ctx.Writer.Size()
		}

		time2 := time.Since(time1)
		fmt.Printf("\nVerb: %s\n Time: %v\n URL: %v\n Size: %d\n Time tx: %v\n", verb, time1, url, size, time2)
		ctx.Next()
	}
}
