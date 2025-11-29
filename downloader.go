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

	file, err := createEmptyFile(d.Destination, fileSize)
	if err != nil {
		return err
	}
	defer file.Close()

	// func caculateChuncks(fileSize, ChunkSize) []Chunk

	chunk := Chunck{
		size:   min(fileSize, 1_048_576),
		offset: 0,
	}

	err = Download(d.Url, chunk, file)
	if err != nil {
		return err
	}

	return nil
}
