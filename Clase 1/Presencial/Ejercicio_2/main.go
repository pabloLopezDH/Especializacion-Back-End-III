package main

import "fmt"

/*
Ejercicio 2 - A qué mes corresponde
Realizar una aplicación que contenga una variable con el número del mes.
Según el número, imprimir el mes que corresponda en texto.
¿Se te ocurre si se puede resolver de más de una manera? ¿Cuál elegirías y por qué?
*/

func main() {

	var meses = []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Obtubre", "Noviembre", "Diciembre"}

	mes := 13

	if mes < 1 || mes > 12 {
		fmt.Println("Error fuera de mes.")
	}

	if mes >= 1 && mes <= 12 {
		for i := 0; i <= len(meses); i++ {
			if i == mes {
				fmt.Println(meses[i-1])
			}
		}
	}

	if mes >= 1 && mes <= 12 {
		fmt.Println(meses[mes-1])
	}

}
