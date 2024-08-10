package qweathersdkgo

import (
	"encoding/json"
	"net/http"
	"net/url"
)

var key = "xxxxxx"

type Client struct {
	APIKey            string
	AirQualityBetaURL string
	WeatherURL        string
	GeoURL            string
	HTTPClient        *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		APIKey:            apiKey,
		AirQualityBetaURL: "https://devapi.qweather.com/airquality/v1/now",
		WeatherURL:        "https://devapi.qweather.com/v7",
		GeoURL:            "https://geoapi.qweather.com/v2",
		HTTPClient:        &http.Client{},
	}
}

func (c *Client) sendRequest(method, endpoint string, params url.Values, v interface{}) error {
	req, err := http.NewRequest(method, endpoint+"?"+params.Encode(), nil)
	if err != nil {
		return err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(v)
}
