package main

import (
	"errors"
	"fmt"
	"log"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func main() {
	minFun, err := operacion(minimum)
	if err != nil {
		log.Fatal(err)
	}

	maxFun, err := operacion(maximum)
	if err != nil {
		log.Fatal(err)
	}

	promFun, err := operacion(average)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("El Minimo es: ", minFun(3, 5, 1, 8, 2, 7, 4))
	fmt.Println("El Maximo es: ", maxFun(3, 5, 1, 8, 2, 7, 4))
	fmt.Println("El Promedio es: ", promFun(3, 5, 1, 8, 2, 7, 4))

	// errorFunc, err := operacion("errorFunc")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("El error es: ", errorFunc(3, 5, 1, 8, 2, 7, 4))

	min, max, avg, err := operacionMulti()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("El minimo es %d | El maximo es: %d | El promedio es: %d",
		min(3, 5, 1, 8, 2, 7, 4),
		max(3, 5, 1, 8, 2, 7, 4),
		avg(3, 5, 1, 8, 2, 7, 4))
}

func calculoMinimo(values ...int) int {
	min := values[0]
	for _, val := range values {
		if val < min {
			min = val
		}
	}
	return min
}

func calculoMaximo(values ...int) int {
	var max int
	for _, val := range values {
		if val > max {
			max = val
		}
	}
	return max
}

func calculoPromedio(values ...int) int {
	sum := 0
	for _, val := range values {
		sum += val
	}
	return sum / len(values)
}

func operacion(calculo string) (func(...int) int, error) {
	switch calculo {
	case minimum:
		return calculoMinimo, nil
	case maximum:
		return calculoMaximo, nil
	case average:
		return calculoPromedio, nil
	default:
		return nil, errors.New("El calculo no exite.")
	}
}

func operacionMulti() (func(...int) int, func(...int) int, func(...int) int, error) {
	return calculoMinimo, calculoMaximo, calculoPromedio, nil
}
