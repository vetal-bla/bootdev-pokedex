package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (c *Client) ExploreLocations(locationName string) (ExploreLocation, error) {
	url := baseURL + "/location-area/" + locationName

	val, ok := c.cache.Get(url)
	if ok {
		exploreLocation := ExploreLocation{}
		err := json.Unmarshal(val, &exploreLocation)
		if err != nil {
			return ExploreLocation{}, nil
		}
		return exploreLocation, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
		return ExploreLocation{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return ExploreLocation{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ExploreLocation{}, err
	}

	if res.StatusCode > 299 {
		log.Fatalf("Responce not ok: %d", res.StatusCode)
	}

	exploreLocation := ExploreLocation{}

	if err := json.Unmarshal(body, &exploreLocation); err != nil {
		log.Fatal("Cant parse json to struct")
		return ExploreLocation{}, err
	}

	c.cache.Add(url, body)

	return exploreLocation, nil
}
