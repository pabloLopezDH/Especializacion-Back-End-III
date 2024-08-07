package main

import "fmt"

/*
Repetir el proceso de la ejercitación realizada en clase, pero ahora implementando fmt.Errorf() para que el mensaje de error reciba por parámetro el valor
 de salary indicando que no alcanza el mínimo imponible. El mensaje mostrado por consola deberá decir: “Error: el mínimo imponible es de 150.000 y el salario ingresado
  es de: [salary]” (siendo [salary] el valor de tipo int pasado por parámetro).

*/

type errSalary struct {
	message string
}

func (e errSalary) Error() string {
	return e.message
}

func main() {

	var salary = 100000
	err := impuesto(salary)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("\n%d Debe pagar impuesto", salary)
	}

}

func impuesto(salary int) error {

	var minImp = 150000

	if salary <= minImp {
		return fmt.Errorf("Error: el mínimo imponible es de %d y el salario ingresado es de: %d", minImp, salary)
	}
	return nil
}
