package utils

import (
	"encoding/json"
	"gonews/models"
	"io/ioutil"
	"log"
	"net/http"
)

func Get(url string) models.Response {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error cannot get data from url")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	var data models.Response

	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err.Error())
	}

	return data
}
