package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationDetail(location string) (RespLocationDetail, error) {
	url := baseURL + "/location-area/" + location

	if val, ok := c.cache.Get(url); ok {
		res := RespLocationDetail{}
		err := json.Unmarshal(val, &res)
		if err != nil {
			return RespLocationDetail{}, err
		}

		return res, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationDetail{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationDetail{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespLocationDetail{}, err
	}

	locationDetail := RespLocationDetail{}
	err = json.Unmarshal(data, &locationDetail)
	if err != nil {
		return RespLocationDetail{}, err
	}

	c.cache.Add(url, data)
	return locationDetail, nil
}
