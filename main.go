package main

func main() {
	// flags
	url, dist, workers := "https://d37ci6vzurychx.cloudfront.net/trip-data/yellow_tripdata_2018-05.parquet", "/Users/medhatmohammed/Documents/goprojects/rangerover/data", 10

	downloader := Downloader{
		Url:         url,
		Destination: dist,
		Workers:     workers,
	}

	downloader.Start()
}
