package fetcher

import (
	"fmt"
	"io"
	"net/http"
)

type FetchResult struct {
	StatusCode int
	ByteCount  string
}

func FetchUrl(url string) (FetchResult, error) {
	resp, err := http.Get(url)
	if err != nil {
		return FetchResult{}, fmt.Errorf("failed to fetch this URL %s: %w", url, err)
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return FetchResult{}, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return FetchResult{
			StatusCode: resp.StatusCode,
			ByteCount:  string(bytes),
		}, fmt.Errorf("unexpected status %d", resp.StatusCode)
	}

	return FetchResult{
		StatusCode: resp.StatusCode,
		ByteCount:  string(bytes),
	}, nil
}