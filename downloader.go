package main

import (
	"fmt"
	"log"
	"net/http"
)

type Downloader struct {
	Url         string
	Destination string
	Workers     int
}

func (d Downloader) Start() error {
	rangesHeaders, err := GetRangeHeaders(d.Url)
	if err != nil {
		return fmt.Errorf("error getting the range headers from the %s\n", d.Url)
	}
	fileSize := rangesHeaders.ContentLength
	err = createEmptyFile(d.Destination, fileSize)
	if err != nil {
		return fmt.Errorf("error getting the range headers from the %s\n", d.Url)
	}
	return nil
}

func GetRangeHeaders(url string) (RangesHeaders, error) {
	resp, err := http.Head(url)
	if err != nil {
		log.Fatalf("error HEADing the %s: %#v\n", url, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Status code is not okay it is %d\n", resp.StatusCode)
	}
	rangesHeaders := RangesHeaders{
		resp.Header.Get("Accept-Ranges"),
		resp.ContentLength,
		resp.Header.Get("ETag"),
	}

	return rangesHeaders, nil
}
