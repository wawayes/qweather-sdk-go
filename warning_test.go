package qweathersdkgo

import (
	"fmt"
	"testing"
)

func TestGetWarningWeather(t *testing.T) {
	c := NewClient(key)
	resp, err := c.GetWarningWeather("101010100")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	fmt.Println(resp)
}

func TestGetWarningList(t *testing.T) {
	c := NewClient(key)
	resp, err := c.GetWarningList()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	fmt.Println(resp)
}
