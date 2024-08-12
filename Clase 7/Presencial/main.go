package main

import (
	"fmt"
	"time"
)

/*
Una empresa de finanzas procesa órdenes de préstamos y órdenes de devoluciones en una proporción
de dos a uno (es decir, por cada devolución dan dos préstamos). Aproximadamente, se procesa una
devolución por segundo, y nos han solicitado que programemos un modelo de procesamiento concurrente
 que denote este modelo.
Se requiere que corran al menos dos goroutines de forma concurrente que procesen estas órdenes.
Donde además se muestre por la consola una devolución por cada dos préstamos hasta que se
apriete cualquier botón y se acabe con la ejecución del programa. El output esperado debería ser similar a este:
devolucion procesada
prestamo procesado
prestamo procesado
devolucion procesada
prestamo procesado
prestamo procesado
*/

func main() {
	devolucion := make(chan string)
	prestamo := make(chan string)

	go func() {
		for {
			time.Sleep(time.Second)
			//fmt.Println("devolucion procesada")
			devolucion <- "devolucion procesada"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second / 2)
			//fmt.Println("prestamo procesado")
			prestamo <- "prestamo procesado"
		}
	}()

	/*go func() {
		for dev := range devolucion {
			fmt.Println(dev)
		}
	}()

	go func() {
		for pres := range prestamo {
			fmt.Println(pres)
		}
	}()
	*/
	go func() {
		for {
			select {
			case msg := <-devolucion:
				println(msg)
			case msg := <-prestamo:
				println(msg)
			}
		}
	}()

	fmt.Scanln()

}
