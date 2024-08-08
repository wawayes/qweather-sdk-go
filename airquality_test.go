package qweathersdkgo

import "testing"

func TestGetAirQuality(t *testing.T) {
	c := NewClient(key)
	r, err := c.GetAirQuality("101010100")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	t.Log(r)
}
