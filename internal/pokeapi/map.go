package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type AreaStruct struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas(url string) (AreaStruct, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return AreaStruct{}, err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Responce not ok: %d", res.StatusCode)
	}

	if err != nil {
		log.Fatal(err)
		return AreaStruct{}, err
	}

	var Area AreaStruct

	if err := json.Unmarshal(body, &Area); err != nil {
		log.Fatal("Cant parse json to struct")
		return AreaStruct{}, err
	}

	return Area, nil
}
