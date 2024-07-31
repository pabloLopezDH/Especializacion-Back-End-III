package main

import "fmt"

func main() {
	var a int
	a = 10
	fmt.Println("Valor de a es: ", a, " el espacio de memoria: ", &a)
	calc(&a)
	fmt.Println("\nCambio de valor a a: ", a, " el espacio de memoria: ", &a)

	var ptr *int
	ptr = &a
	fmt.Println(ptr)
	*ptr = 30
	fmt.Println("\nPuntero ptr: ", ptr, " el valor es: ", *ptr)
	fmt.Println("\nImprimo a final", a)
}

func calc(a *int) {
	fmt.Println("Lo que me llega: ", a, " y el valor desreferenciado: ", *a)
	*a = 20
	fmt.Println("Valor de a: ", a, " y el valor desreferenciado: ", *a)
}
