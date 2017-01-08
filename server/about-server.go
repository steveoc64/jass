package main

import "github.com/labstack/echo"

func aboutTracker(c echo.Context) error {
	req := c.Request()
	printLog(c, "direct about, ref=", req.Referer())
	return c.File("public/index.html")
}

func contactTracker(c echo.Context) error {
	req := c.Request()
	printLog(c, "direct contact, ref=", req.Referer())
	return c.File("public/index.html")
}
