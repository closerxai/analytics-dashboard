package dashboard

import (
	"encoding/json"
	"net/http"
	"time"
)

type HealthResponse struct {
	Status string `json:"status"`
}

// Shared HTTP client with timeout
var client = &http.Client{
	Timeout: 5 * time.Second,
}

// Reusable function to check ANY external health endpoint
func CheckExternalHealth(url string) (bool, error) {
	resp, err := client.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return false, nil
	}

	var body HealthResponse
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return false, err
	}

	return body.Status == "ok", nil
}
