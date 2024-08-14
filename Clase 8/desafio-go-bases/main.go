package main

import (
	"desafio-ctd/internal/tickets"
	"desafio-ctd/test"
	"fmt"
)

func main() {
	total, _ := tickets.GetTotalTickets("Brazil")

	test.Hello()

	fmt.Println(total)
}
