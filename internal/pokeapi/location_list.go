package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


func (c *Client) ListLocations(pageUrl *string) (LocationList, error) {
	url := baseURL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	if data, found := c.cache.Get(url); found {
		var locationList LocationList
		err := json.Unmarshal(data, &locationList)
		if err != nil {
			return LocationList{}, err
		}
		return locationList, nil
		}
	

	req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationList{}, err
		}

	res, err := c.httpClient.Do(req)
		if err != nil {
			return LocationList{}, err
		}

	defer res.Body.Close() 

	body, err := io.ReadAll(res.Body)
		if err != nil {
			return LocationList{}, err
		}

	var locationList LocationList
	err = json.Unmarshal(body, &locationList)
		if err != nil {
			return LocationList{}, err
		}

	c.cache.Add(url, body)
	return locationList, nil
}


func (c *Client) GetLocation(name string) (Location, error) {
	url := baseURL + "/location-area" + "/" + name

	if data, found := c.cache.Get(url); found {
		var location Location
		err := json.Unmarshal(data, &location)
		if err != nil {
			return Location{}, err
		}
		return location, nil
		}
	
	req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Location{}, err
		}

	res, err := c.httpClient.Do(req)
		if err != nil {
			return Location{}, err
		}

	defer res.Body.Close()
	
	if res.StatusCode != http.StatusOK {
		return Location{}, fmt.Errorf("unexpected status %d", res.StatusCode)
	}


	body, err := io.ReadAll(res.Body)
		if err != nil {
			return Location{}, err
		}

	var location Location
	err = json.Unmarshal(body, &location)
		if err != nil {
			return Location{}, err
		}

	c.cache.Add(url, body)
	return location, nil
}
