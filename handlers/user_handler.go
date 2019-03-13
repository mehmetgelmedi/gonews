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

	return c.JSON(http.StatusCreated, echo.Map{
		"UserId": user.ID,
		"Result": registrationCheckHelper(user),
	})
}

func registrationCheckHelper(user *models.User) string {
	if user.ID > 0 {
		return "Account successfully created. Thank you for your registration!"
	 } else { 
		return "The username " + user.Username + " already exists!"
	}
}
