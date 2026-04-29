package server

import (
	"io"
	"net/http"
)

var client = &http.Client{}

type FetchConfig struct {
	Method string
	Token  string
}

func Fetch(url string, config FetchConfig) (string, error) {
	req, err := http.NewRequest(config.Method, url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.Token)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	return string(body), nil
}
