package qweathersdkgo

import (
	"fmt"
	"testing"
)

func TestGetGridCurrentWeather(t *testing.T) {
	client := NewClient(key)
	r, err := client.GetGridCurrentWeather("116.41,39.92")
	if err != nil {
		fmt.Println("err", err)
		t.Error(err)
	}
	fmt.Println(r)
}

func TestGetGridDailyWeather(t *testing.T) {
	client := NewClient(key)
	r, err := client.GetGridDailyWeather("116.4123,39.9232", 3)
	if err != nil {
		fmt.Println("err", err)
		t.Error(err)
	}
	fmt.Println(r)
}

func TestGetGridHourlyWeather(t *testing.T) {
	client := NewClient(key)
	r, err := client.GetGridHourlyWeather("116.41,39.92", 24)
	if err != nil {
		fmt.Println("err", err)
		t.Error(err)
	}
	fmt.Println(r)
}
