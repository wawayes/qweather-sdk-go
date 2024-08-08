package qweathersdkgo

import "testing"

func TestGetIndicesWeather(t *testing.T) {
	c := NewClient(key)
	resp, err := c.GetIndicesWeather("3,8", "101010100", 1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	t.Log(resp)
}
