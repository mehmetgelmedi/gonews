package handlers

import (
	"github.com/labstack/echo"
)

// Handler
func AllHello(c echo.Context) error {
	return c.JSON(200, "ALL News")
}

func FindHello(c echo.Context) error {
	return c.JSON(200, "FIND News "+c.Param("id"))
}

func CreateHello(c echo.Context) error {
	return c.JSON(200, "CREATE News")
}

func UpdateHello(c echo.Context) error {
	return c.JSON(200, "UPDATE News")
}

func DeleteHello(c echo.Context) error {
	return c.JSON(200, "DELETE News")
}
