package main

import (
	"fmt"
	"net/http"
	"time"
)

type RangesHeaders struct {
	AcceptRanges  string
	ContentLength int64
	ETag          string
}

type Chunck struct {
	offset int64
	size   int64
}

func Download(url string, chunk Chunck, dist string) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("error creating a new GET/ request for url: %s\n", url)
	}

	rangeVal := fmt.Sprintf("bytes=%d-%d", chunk.size*chunk.offset, chunk.size*(chunk.offset+1))
	req.Header.Set("Range", rangeVal)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending a new GET request %#v for url: %s\n", req, url)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Status code is not okay it is %d\n", resp.StatusCode)
	}
	fmt.Print(resp.Body)
	return nil
}
