package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (c *Client) GetLocationAreas(pageUrl *string) (AreaStruct, error) {
	url := baseURL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
		return AreaStruct{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return AreaStruct{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return AreaStruct{}, err
	}

	if res.StatusCode > 299 {
		log.Fatalf("Responce not ok: %d", res.StatusCode)
	}

	Area := AreaStruct{}

	if err := json.Unmarshal(body, &Area); err != nil {
		log.Fatal("Cant parse json to struct")
		return AreaStruct{}, err
	}

	return Area, nil
}
