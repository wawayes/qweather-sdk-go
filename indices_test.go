package qweathersdkgo

import "testing"

func TestGetIndicesWeather(t *testing.T) {
	c := NewClient(key)
	indicesTypeSlice := []string{"0"}
	resp, err := c.GetIndicesWeather(indicesTypeSlice, "101010100", 1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	t.Log(resp)
}
