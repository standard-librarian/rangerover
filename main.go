package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("creating a sparce file in the download")
	path, size := "/Users/medhatmohammed/Downloads/new.json", int64(3500)
	if err := createEmptyFile(path, size); err != nil {
		log.Fatalf("error creating the file %#v", err)
	}
	fmt.Println("creating done successfully")

}
