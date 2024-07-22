package main

import "fmt"

/*
Ejercicio 1 - Letras de una palabra
La Real Academia Española quiere saber cuántas letras tiene una palabra y
luego tener cada una de las letras por separado para deletrearla.
Crear una aplicación que tenga una variable con la palabra e imprimir la
cantidad de letras que tiene la misma.
Luego, imprimir cada una de las letras.
*/

func main() {

	var palabra string = "palabra"
	//palabra2 := "palabra2"

	//cant := len(palabra)
	var cant int = len(palabra)

	var Ok bool = true
	var flotante float64 = 20.64

	fmt.Print("La palabra ", palabra, " tiene: ", cant, " letras.")
	fmt.Println("La palabra ", palabra, " tiene: ", cant, " letras.")
	fmt.Printf("La palabra %s tiene %d letras. %t %f\n", palabra, cant, Ok, flotante)
	fmt.Printf("Boleano %t - flotante %f\n", Ok, flotante)
	fmt.Println()

	for i := 0; i < cant; i++ {
		fmt.Println(string(palabra[i]))
	}

}
