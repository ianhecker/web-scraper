package scrape

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func Get(urlStr string, pageNumber int) ([]byte, int, error) {
	URL, err := url.Parse(urlStr)
	if err != nil {
		return nil, 0, fmt.Errorf("error parsing url: %w", err)
	}
	q := URL.Query()
	q.Set("pageno", fmt.Sprintf("%d", pageNumber))
	URL.RawQuery = q.Encode()

	resp, err := http.Get(URL.String())
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
