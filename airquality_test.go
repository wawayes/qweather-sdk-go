package qweathersdkgo

import "testing"

func TestGetAirQuality(t *testing.T) {
	c := NewClient(key)
	r, err := c.GetAirQuality("116.4123,39.9232")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	t.Log(r)
}
