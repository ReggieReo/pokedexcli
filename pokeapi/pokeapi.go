package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Location struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocation(url string) (next, prvious string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	allLocation := Location{}
	json.Unmarshal(body, &allLocation)
	printLocation(&allLocation)
	return allLocation.Next, allLocation.Previous
}

func printLocation(location *Location) {
	for _, city := range location.Results {
		fmt.Println(city.Name)
	}
}
