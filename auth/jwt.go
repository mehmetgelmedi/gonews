package auth

import (
	"gonews/models"
	"net/http"
	"time"
	"os"
	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/middleware"
)

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Throws unauthorized error
	if username != "jon" || password != "shhh!" {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &models.JwtCustomClaims{
		"Jon Snow",
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func Accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to Gonews API powered by NewsAPI")
}

func Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*models.JwtCustomClaims)
	//name := claims.Name
	return c.JSON(http.StatusOK, claims)
}

func Config() middleware.JWTConfig{
	secret := os.Getenv("SECRET")
	config := middleware.JWTConfig{
		Claims:     &models.JwtCustomClaims{},
		SigningKey: []byte(secret),
	}
	return config
}