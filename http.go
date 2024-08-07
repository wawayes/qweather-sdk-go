package qweathersdkgo

import (
	"io/ioutil"
	"net/http"
)

// HttpGet sends an HTTP GET request and returns the response body and error.
func HttpGet(urlPath string) ([]byte, error) {
	resp, err := http.Get(urlPath)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
