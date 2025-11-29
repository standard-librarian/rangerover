package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type PartWriter struct {
	File   *os.File
	Offset int64
}

func (pw *PartWriter) Write(p []byte) (n int, err error) {
	n, err = pw.File.WriteAt(p, pw.Offset)
	pw.Offset += int64(n)
	return n, err
}

// createEmptyFile creates an empty file in given size
func createEmptyFile(path string, size int64) (*os.File, error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	// defer file.Close()
	file.Seek(size-1, io.SeekStart)
	file.Write([]byte{0})
	return file, nil
}

// GetRangeHeaders gets Accept-Ranges, Content-Length, and Etag header using a Head http request
func GetRangeHeaders(url string) (RangesHeaders, error) {
	resp, err := http.Head(url)
	if err != nil {
		return RangesHeaders{}, fmt.Errorf("error heading url:%s\n", url)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return RangesHeaders{}, fmt.Errorf("Status code is not okay it is %d\n", resp.StatusCode)
	}
	rangesHeaders := RangesHeaders{
		resp.Header.Get("Accept-Ranges"),
		resp.ContentLength,
		resp.Header.Get("ETag"),
	}
	return rangesHeaders, nil
}
