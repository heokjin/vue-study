package main

import (
	"github.com/labstack/echo"
)

func Hello(c echo.Context) error {

	return c.HTML(200, "index")
}
