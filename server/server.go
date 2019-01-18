package server

import (
	"gonews/handlers"
	"os"

	"github.com/labstack/echo"
)

func Run() {
	// Echo instance
	e := echo.New()

	// Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	/*** Routes ***/

	// News Routes
	e.File("/", "asset/index.html")
	e.GET("/news", handlers.GetAllNews)
	e.GET("/news/:id", handlers.GetNewsByID)

	// Start server
	serverPort := os.Getenv("SERVER_PORT")
	e.Logger.Fatal(e.Start(":" + serverPort))
}
