package netops

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func Fetch(request *http.Request) (string, error) {
	client := &http.Client{Timeout: 100 * time.Second}

	response, err := client.Do(request)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %w", err)
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return "", fmt.Errorf("request failed with status code: %d", response.StatusCode)
	}

	responseText, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	return string(responseText), nil
}
