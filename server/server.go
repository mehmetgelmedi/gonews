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
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root: "frontend/dist",
	}))
	
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{echo.GET},
	}))

	/*** Routes ***/
	// News Routes
	e.GET("/news", handlers.GetAllNews)
	e.GET("/news/:search", handlers.GetNewsBySearch)

	// Start server
	serverPort := os.Getenv("SERVER_PORT")
	e.Logger.Fatal(e.Start(":" + serverPort))
}
