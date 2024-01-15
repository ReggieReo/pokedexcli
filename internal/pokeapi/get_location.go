package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocation(url *string) (Location, error) {
	endpoint := "/location"
	fullURL := baseURL + endpoint
	// make Location struct for
	locationAreasResp := Location{}
	if url != nil {
		fullURL = *url
	}

	cDat, ok := c.cache.Get(fullURL)
	// if there cache unmashal data and return them
	if ok {
		err := json.Unmarshal(cDat, &locationAreasResp)
		if err != nil {
			return Location{}, err
		}
		return locationAreasResp, nil
	}

	// else make a request normally
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
	c.cache.Add(fullURL, dat)
	err = json.Unmarshal(dat, &locationAreasResp)

	if err != nil {
		return Location{}, err
	}

	return locationAreasResp, nil

}
