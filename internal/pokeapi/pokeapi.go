package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

func (c *Client) GetLocation(url *string) (Location, error) {
	endpoint := "/location"
	fullURL := baseURL + endpoint
	if url != nil {
		fullURL = *url
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return Location{}, fmt.Errorf("bad status code:%v", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	locationAreasResp := Location{}
	err = json.Unmarshal(dat, &locationAreasResp)
	if err != nil {
		return Location{}, err
	}

	return locationAreasResp, nil

}
