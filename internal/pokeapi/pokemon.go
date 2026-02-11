package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (c *Client) PokemonInfo(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	val, ok := c.cache.Get(url)
	if ok {
		pokemon := Pokemon{}
		err := json.Unmarshal(val, &pokemon)
		if err != nil {
			return Pokemon{}, nil
		}
		return pokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
		return Pokemon{}, nil
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, nil
	}

	if res.StatusCode > 299 {
		log.Fatal("Responce not ok: %d", res.StatusCode)
	}

	pokemon := Pokemon{}

	if err := json.Unmarshal(body, &pokemon); err != nil {
		log.Fatal("Cant parse json to struct")
		return Pokemon{}, err
	}

	c.cache.Add(url, body)

	return pokemon, nil
}
