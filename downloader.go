package main

type Downloader struct {
	Url         string
	Destination string
	Workers     int
}

func (d Downloader) Start() error {
	rangesHeaders, err := GetRangeHeaders(d.Url)
	if err != nil {
		return err
	}

	fileSize := rangesHeaders.ContentLength
	if err := createEmptyFile(d.Destination, fileSize); err != nil {
		return err
	}
	chunk := Chunck{
		size:   1024,
		offset: 0,
	}
	// func caculateChuncks(fileSize, ChunkSize) []Chunk
	Download(d.Url, chunk, d.Destination)

	return nil
}
