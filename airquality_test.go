package qweathersdkgo

import "testing"

func TestGetAirQuality(t *testing.T) {
	c := NewClient(key)
	r, err := c.GetAirQuality("10101010")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	t.Log(r)
}
