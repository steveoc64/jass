package main

import (
	"net/http"

	"../shared"
	"github.com/labstack/echo"
)

func getProducts(c echo.Context) error {
	// println("getting items")

	prods := []shared.Product{}
	err := DB.SQL(`select * from product order by name desc`).QueryStructs(&prods)
	if err != nil {
		println(err.Error())
		return c.String(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, prods)
}
