package handlers

import (
	"encoding/json"
	"gonews/models"
	"gonews/utils"
	"os"

	"github.com/labstack/echo"
)

// Handler
func GetAllNews(c echo.Context) error {
	url := "https://newsapi.org/v2/top-headlines?country=us&apiKey=" + os.Getenv("API_KEY")
	//var data models.NewsList
	res := utils.Get(url)

	v, ok := res.(map[string]interface{})
	if !ok {
		return c.JSON(500, "undefined")
	}
	var newss []models.News
	err := json.Unmarshal(v["articles"], &newss)
	//fmt.Println(reflect.TypeOf(v["articles"])) []interface {}
	return c.JSON(200, "")
}

func GetNewsByID(c echo.Context) error {
	return c.JSON(200, "FIND News "+c.Param("id"))
}
