package handlers

import (
	"gonews/utils"
	"os"

	"github.com/labstack/echo"
)

// Handler
func GetAllNews(c echo.Context) error {
	url := "https://newsapi.org/v2/top-headlines?country=us&apiKey=" + os.Getenv("API_KEY")
	//var data models.NewsList
	res := utils.Get(url)

	return c.JSON(200, res.Articles)
}

func GetNewsBySearch(c echo.Context) error {
	search := c.Param("search")
	url := "https://newsapi.org/v2/everything?q=" + search + "&apiKey=" + os.Getenv("API_KEY")
	res := utils.Get(url)
	return c.JSON(200, res.Articles)
}
