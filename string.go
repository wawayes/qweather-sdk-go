package qweathersdkgo

import (
	"net/url"
	"strings"
)

// ConcatURLWithSuffix concatenates a URL with variable number of suffixes.
func ConcatURLWithSuffix(baseURL string, suffixes ...string) (string, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	for _, suffix := range suffixes {
		u.Path = strings.TrimSuffix(u.Path, "/") + "/" + strings.TrimPrefix(suffix, "/")
	}

	return u.String(), nil
}
