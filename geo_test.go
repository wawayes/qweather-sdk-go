package qweathersdkgo

import "testing"

func TestCityLookup(t *testing.T) {
	c := NewClient(key)
	resp, err := c.CityLookup("116.27,40.15", "")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	t.Logf("Response: %+v", resp)
}
