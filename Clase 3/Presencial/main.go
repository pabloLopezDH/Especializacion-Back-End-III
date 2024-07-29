package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Id          int    `json:"id"`
	Name        string `json:"name,omitempty"`
	DateOfBirth string `json:"-"`
}
type Employee struct {
	Id       int    `json:"id"`
	Position string `json:"position"`
	Person   Person `json:"person"`
}

// metodo asociado a employee, estructura que devuelve string
func (e Employee) PrintEmployee() string {
	return fmt.Sprintf("Employee ID: %d\n Name: %s\n Date of Birth %s\n Position: %s", e.Id, e.Person.Name, e.Person.DateOfBirth, e.Position)
}

// metodo asociado a employee, estructura que no asignamos que devuelve - no usa return -
func (e Employee) PrintEmployee2() {
	fmt.Printf("Employee ID: %d\n Name: %s\n Date of Birth %s\n Position: %s", e.Id, e.Person.Name, e.Person.DateOfBirth, e.Position)
}

func main() {
	p1 := Person{1, "Juan", "10/10/2000"}
	e1 := Employee{1, "Gerente", p1}
	//Llamamos al método que retorna string
	fmt.Println(e1.PrintEmployee())
	// Llamamos al método que no le asignamos que devuelve (no usamos return)
	e1.PrintEmployee2()

	employee, err := json.Marshal(e1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	fmt.Println()
	fmt.Println(string(employee))

	fmt.Println()
	fmt.Println()

	var e2 Employee

	//employee2 := "{\"id\":2,\"position\":\"Developer\",\"person\":{\"id\":2,\"name\":\"Mariano\"}}"
	data := `{"Id":1,"Position":"Gerente","Person":{"Id":1,"Name":"Juan","DateOfBirth":"10/10/2000"}}`

	if err2 := json.Unmarshal([]byte(data), &e2); err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(e2.PrintEmployee())

	fmt.Println("algo \"entre\" comillas")
}
