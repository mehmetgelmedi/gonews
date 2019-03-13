package handlers

import (
	"gonews/models"
	"gonews/repositories"
	"net/http"

	"github.com/labstack/echo"
)

func Register(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		panic(err)
	}

	user = repositories.CreateUser(user)

	return c.JSON(http.StatusCreated, user.ID)
}
