package main

import (
	"flag"
	"log"
)

var (
	url                string
	dist               string
	numberOfGoroutines int
	chunckSize         int
	numberOfRetries    int
)

func init() {
	flag.StringVar(&url, "url", "", "url of the file needed to be downloaded")
	flag.StringVar(&dist, "dist", "", "the destination of the file to be downloaded")
	flag.IntVar(&numberOfGoroutines, "n", 1, "parameter to limit the number of downloading goroutines")
	flag.IntVar(&chunckSize, "s", 1_048_576, "parameter to set the max chunk size")
	flag.IntVar(&numberOfRetries, "r", 3, "parameter to control number of retries")

	flag.Parse()
}

func main() {

	d := Downloader{
		Url:         url,
		Destination: dist,
		Workers:     numberOfGoroutines,
	}

	err := d.Start()
	if err != nil {
		log.Fatal(err)
	}
}
