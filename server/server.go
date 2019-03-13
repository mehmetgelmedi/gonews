package server

import (
	"gonews/auth"
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
		AllowMethods: []string{echo.GET, echo.POST},
	}))

	/*** Routes ***/
	// News Routes
	e.POST("/register", handlers.Register)
	e.POST("/login", auth.Login)
	e.GET("/", auth.Accessible)

	r := e.Group("/api")

	config := auth.Config()

	r.Use(middleware.JWTWithConfig(config))
	r.GET("", auth.Restricted)

	r.GET("/news", handlers.GetAllNews)
	r.GET("/news/:search", handlers.GetNewsBySearch)

	// Start server
	serverPort := os.Getenv("SERVER_PORT")
	e.Logger.Fatal(e.Start(":" + serverPort))
}
