package main

import "fmt"

/*
Una empresa de inteligencia artificial necesita tener una funcionalidad para crear matrices en Go.
 Para ello requieren una estructura Matrix que tenga los métodos:
Set: recibe una matriz de flotantes e inicializa los valores en la estructura Matrix.
Print: imprime por pantalla la matriz de una forma más visible (con los saltos de línea entre filas).
La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión
 del ancho, si es cuadrática y cuál es el valor máximo.
*/

type Matrix struct {
	valores []float64
	alto    int
	ancho   int
	max     float64
	cuadr   bool
}

func main() {
	m := Matrix{
		valores: []float64{},
		alto:    3,
		ancho:   4,
	}
	fmt.Println(m)
	valores := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	m.Set(valores...)
	fmt.Println(m)

	m.Print()
	m.Max()

	fmt.Println("Es cuadratica? ", m.cuadrática())

	fmt.Println("\n", m)
	fmt.Println("\nEspacio de memoria", m)
}

func (m *Matrix) Set(valores ...float64) {
	if len(valores) != m.alto*m.ancho {
		fmt.Println("La cantidad de valores no coincide con el alto y ancho")
		return
	}
	m.valores = append(m.valores, valores...)
}

func (m Matrix) Print() {
	for fila := 0; fila < m.alto; fila++ {
		fmt.Println(m.valores[fila*m.ancho : fila*m.ancho+m.ancho])
	}
	//fila = 0 - [0:3]
	//fila = 1 - [3:6]
	//fila = 2 - [6:9]
}

func (m *Matrix) Max() {
	max := m.valores[0]
	for _, elemento := range m.valores {
		if elemento > max {
			max = elemento
		}
	}
	m.max = max
	fmt.Println("El maximo es: ", max)
}

func (m *Matrix) cuadrática() bool {
	var esCuadr bool
	if m.alto == m.ancho && m.alto != 0 {
		esCuadr = true
		return esCuadr
	}
	m.cuadr = false
	return esCuadr
}
