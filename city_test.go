package qweathersdkgo

import (
	"fmt"
	"testing"
)

func TestGetCurrentWeather(t *testing.T) {
	client := NewClient(key)
	r, err := client.GetCurrentWeather("101010100")
	if err != nil {
		fmt.Println("err", err)
		t.Error(err)
	}
	fmt.Println(r)
}

func TestGetWeatherForcast(t *testing.T) {
	client := NewClient(key)
	r, err := client.GetDailyForecast("101010100", 3)
	if err != nil {
		fmt.Println("err", err)
		t.Error(err)
	}
	fmt.Println(r)
}

func TestGetHourlyWeather(t *testing.T) {
	client := NewClient(key)
	r, err := client.GetHourlyWeather("101010100", 24)
	if err != nil {
		fmt.Println("err", err)
		t.Error(err)
	}
	fmt.Println(r)
}
