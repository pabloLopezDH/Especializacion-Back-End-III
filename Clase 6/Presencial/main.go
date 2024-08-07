package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
La empresa encargada de vender los productos de limpieza ahora necesita leer el archivo almacenado.
Para esto requiere que se imprima por pantalla mostrando los valores tabulados, con un título
(tabulado a la izquierda para el ID y a la derecha para el Precio y Cantidad), el precio, la cantidad
 y, abajo del precio, se debe visualizar el total (sumando precio por cantidad).
*/

func main() {

	res, err := os.ReadFile("./data.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	//fmt.Println(string(res))

	data := strings.Split(string(res), ";")
	//fmt.Println(data)

	var total float64

	for i := 0; i < len(data)-1; i++ {

		var line = strings.Split(string(data[i]), ",")
		fmt.Println(line[0], "\t\t", line[1], "\t\t", line[2])

		if i != 0 {
			precio, err := strconv.ParseFloat(line[1], 64)
			if err != nil {
				println("el precio no se pudo parsear")
			}

			cantidad, err := strconv.ParseFloat(line[2], 64)
			if err != nil {
				println("la cantidad no se pudo parsear")
			}

			totalProducto := precio * cantidad
			total += totalProducto

		}
		//w := tabwriter.NewWriter(os.Stdout, 10, 1, 1, ' ', tabwriter.Debug)
		//fmt.Fprintf(w, "%s\t%s\t%s\t\n", line[0], line[1], line[2])
		// fmt.Printf("%s\t\t%s\t\t%s\t\t\n", line[0], line[1], line[2])

		// for i := 0; i < len(data)-1; i++ {
		// 	fmt.Printf("%s\t\t", line[i])
		// 	if i == len(line)-1 {
		// 		fmt.Print("\n")
		// 	}
		// }

	}
	fmt.Printf("Total:\t\t%.2f", total)

}
