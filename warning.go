package qweathersdkgo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WarningResponse struct {
	Code       string `json:"code"`
	UpdateTime string `json:"updateTime"`
	FxLink     string `json:"fxLink"`
	Warning    []struct {
		ID        string `json:"id"`
		Sender    string `json:"sender"`
		PubTime   string `json:"pubTime"`
		Title     string `json:"title"`
		StartTime string `json:"startTime"`
		EndTime   string `json:"endTime"`
		Status    string `json:"status"`
		Level     string `json:"level"`
		Type      string `json:"type"`
		TypeName  string `json:"typeName"`
		Text      string `json:"text"`
		Related   string `json:"related"`
	} `json:"warning"`
}

// GetWarning retrieves the current weather warnings for a location
func (c *Client) GetWarning(location string, lang string) (*WarningResponse, error) {
	params := map[string]string{
		"location": location,
	}
	if lang != "" {
		params["lang"] = lang
	}

	url, err := c.buildURL("/warning/now", params)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	var warningResp WarningResponse
	if err := json.NewDecoder(resp.Body).Decode(&warningResp); err != nil {
		return nil, err
	}

	return &warningResp, nil
}
