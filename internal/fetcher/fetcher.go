package fetcher

import (
	"fmt"
	"io"
	"net/http"
)

type FetchResult struct {
	StatusCode int
	ByteCount  int64
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

	return FetchResult{
		StatusCode: resp.StatusCode ,
		ByteCount: int64(len(bytes)),
	}, nil
}