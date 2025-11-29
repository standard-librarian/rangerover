package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type RangesHeaders struct {
	AcceptRanges  string
	ContentLength int64
	ETag          string
}

type Chunk struct {
	offset int64
	size   int64
}

func Download(url string, chunk Chunk, file *os.File) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("error creating GET request for url %s: %w", url, err)
	}

	rangeVal := fmt.Sprintf("bytes=%d-%d", chunk.size*chunk.offset, chunk.size*chunk.offset+chunk.size-1)
	req.Header.Set("Range", rangeVal)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending GET request to %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusPartialContent {
		return fmt.Errorf("Status code is not okay it is %d\n", resp.StatusCode)
	}

	pw := PartWriter{
		File:   file,
		Offset: chunk.offset,
	}

	if _, err := io.Copy(&pw, resp.Body); err != nil {
		log.Fatal(err)
	}

	return nil
}
