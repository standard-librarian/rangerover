package main

import (
	"errors"
	"flag"
	"log"
)

func parseFlags() (*Config, error) {
	cfg := &Config{}
	flag.StringVar(&cfg.Url, "url", "", "URL of the file to download")
	flag.StringVar(&cfg.Dest, "dist", "", "destination path")
	flag.IntVar(&cfg.Workers, "n", 1, "number of parallel workers")
	flag.Int64Var(&cfg.ChunkSize, "s", 1<<20, "chunk size in bytes")
	flag.IntVar(&cfg.MaxRetries, "r", 3, "max retry attempts")
	flag.Parse()

	if cfg.Url == "" || cfg.Dest == "" {
		return nil, errors.New("url and dist are required")
	}
	return cfg, nil
}

func main() {
	cfg, err := parseFlags()
	if err != nil {
		log.Fatal(err)
	}

	d, err := NewDownloader(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	if err := d.Start(); err != nil {
		log.Fatal(err)
	}
}
