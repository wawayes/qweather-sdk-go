package qweathersdkgo

import (
	"fmt"
	"testing"
)

func TestConcatURLWithSuffix(t *testing.T) {
	baseURL := "https://example.com"
	suffixes := []string{"path1", "path2", "path3"}

	expectedResult := "https://example.com/path1/path2/path3"

	result, err := ConcatURLWithSuffix(baseURL, suffixes...)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if result != expectedResult {
		t.Errorf("Expected %s, but got %s", expectedResult, result)
	}

	fmt.Println(result)
}
