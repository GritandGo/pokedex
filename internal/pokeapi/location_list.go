package pokeapi
import (
	"io"
	"net/http"
	"encoding/json"
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
	