package main

import (
	"flag"
	"fmt"
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
	flag.StringVar(&dist, "dist", "", "the distination of the file to be downloaded")
	flag.IntVar(&numberOfGoroutines, "n", 1, "parameter to limit the number of downloading goroutines")
	flag.IntVar(&chunckSize, "s", 1_048_576, "parameter to set the max chunk size")
	flag.IntVar(&numberOfRetries, "r", 3, "parameter to control number of retries")

	flag.Parse()
}

func main() {

	downloader := Downloader{
		Url:         url,
		Destination: dist,
		Workers:     numberOfGoroutines,
	}

	fmt.Println(downloader.Start())
}
