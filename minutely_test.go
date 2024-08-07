package qweathersdkgo

import (
	"fmt"
	"testing"
)

func TestGetMinutelyPrecipitation(t *testing.T) {
	client := NewClient(key)
	r, err := client.GetMinutelyPrecipitation("116.41,39.92")
	if err != nil {
		fmt.Println("err", err)
		t.Error(err)
	}
	fmt.Println(r)
}
