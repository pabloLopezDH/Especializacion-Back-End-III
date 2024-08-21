package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id          int
	Name        string
	Quantity    int
	CodeValue   string
	IsPublished bool
	Expiration  string
	Price       float64
}

var productList = []Product{}

func main() {

	loadProducts("./products.json", &productList)
	router := gin.Default()

	router.GET("/products/search", searchProduct())
	router.Run()
}

func searchProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		price := ctx.Query("priceGt") //asdds

		priceGt, err := strconv.ParseFloat(price, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "El valor ingresado no se pudo parsear."})
			return
		}

		list := []Product{}

		for _, product := range productList {
			if product.Price > priceGt {
				list = append(list, product)
			}
		}

		if len(list) > 0 {
			ctx.JSON(http.StatusOK, list)
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "No hay productos mayores al precio ingresado.",
			})
		}

	}
}

func loadProducts(path string, list *[]Product) {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
		// err := recover()
		// if err != nil {
		// 	log.Fatal(err)
		// }
	}()

	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(file, &list)
	if err != nil {
		panic(err)
	}
}
