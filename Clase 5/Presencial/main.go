package main

import (
	"errors"
	"fmt"
)

/*
En la función main, definir una variable llamada salary y asignarle un valor de tipo “int”.
 Crear un error con errors.New() con el mensaje “salario muy bajo" y lanzarlo en caso de
 que salary sea menor o igual a 10.000. La validación debe ser hecha con la función Is()
 dentro del main.*/

var errSalary = errors.New("salario muy bajo")

func main() {

	var salary = 9000
	err := esSalarioBajo(salary)
	if errors.Is(err, errSalary) {
		fmt.Println(err)
	}

}

func esSalarioBajo(salary int) error {
	if salary <= 10000 {
		return errSalary
	}
	return nil
}
