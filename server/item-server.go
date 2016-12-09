package main

import (
	"fmt"
	"net/http"

	"../shared"
	"github.com/labstack/echo"
)

func getItems(c echo.Context) error {
	println("getting items")

	items := []shared.Item{}

	items = append(items, shared.Item{
		SKU:   "abc123",
		Price: 123.45,
		Name:  "200ml Jass Perfume",
		Descr: "A bottle of the worlds finest perfume",
		Image: "img/items/jass_product.jpg",
	})

	fmt.Printf("the items array is %v\n", items)

	return c.JSON(http.StatusOK, items)

}
