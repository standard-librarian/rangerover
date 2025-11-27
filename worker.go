package main

type RangesHeaders struct {
	AcceptRanges  string
	ContentLength int64
	ETag          string
}

type Chunck struct {
	id   int64
	size string
	data []byte
}
