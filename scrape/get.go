package scrape

import (
	"fmt"
	"io"
	"net/http"
)

func Get(url string) ([]byte, int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, 0, fmt.Errorf("error getting: %w", err)
	}
	defer resp.Body.Close()

	status := resp.StatusCode
	if status != http.StatusOK {
		return nil, status, fmt.Errorf("unexpected status code: %d", status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("error reading body: %w", err)
	}
	return body, status, nil
}
