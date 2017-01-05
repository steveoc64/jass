package main

import (
	// "fmt"

	"net/http"
	"strconv"

	"../shared"
	"github.com/labstack/echo"
)

func getBlogs(c echo.Context) error {

	printLog(c, "getBlogs")

	blogs := []shared.Blog{}

	err := DB.SQL(`select * from blog order by post_order desc`).QueryStructs(&blogs)
	if err != nil {
		println(err.Error())
		return c.String(http.StatusNotFound, err.Error())
	}
	// fmt.Printf("the blogs array is %v\n", blogs)

	return c.JSON(http.StatusOK, blogs)
}

func blogTracker(c echo.Context) error {
	req := c.Request()
	printLog(c, "direct blog, ref=", req.Referer())
	return c.File("public/index.html")
}

func blogIDTracker(c echo.Context) error {
	req := c.Request()
	id, _ := strconv.Atoi(c.Param("id"))
	printLog(c, "direct blog article, id=", id, "ref=", req.Referer())
	return c.File("public/index.html")
}
