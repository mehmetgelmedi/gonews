package server

import (
	"gonews/handlers"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Run() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	/*** Routes ***/

	// Hello Routes
	e.GET("/hello", handlers.AllHello)
	e.GET("/hello/:id", handlers.FindHello)
	e.POST("/hello", handlers.CreateHello)
	e.PUT("/hello", handlers.UpdateHello)
	e.DELETE("/hello", handlers.DeleteHello)

	// Start server
	serverPort := os.Getenv("SERVER_PORT")
	e.Logger.Fatal(e.Start(":" + serverPort))
}
