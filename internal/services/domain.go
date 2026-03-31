package services

import (
	"io"
	"net/http"
)

type Domain struct{}

func (d *Domain) Fetch(domain string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://edsail.com/api/domain/"+domain, nil)
	if err != nil {
		return "", err
	}

	// Add the Bearer token here
	req.Header.Set("Authorization", "Bearer 147852")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	return string(body), nil
}
