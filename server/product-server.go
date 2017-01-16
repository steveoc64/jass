package main

import (
	"net/http"

	"../shared"
	"github.com/labstack/echo"
)

func getCategory(c echo.Context) error {
	data := []shared.Category{}
	err := DB.SQL(`select * from category order by id`).QueryStructs(&data)
	if err != nil {
		println(err.Error())
		return c.String(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, data)
}

func getProducts(c echo.Context) error {
	data := []shared.Product{}
	err := DB.SQL(`select * from product order by cat_id,price`).QueryStructs(&data)
	if err != nil {
		println(err.Error())
		return c.String(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, data)
}
