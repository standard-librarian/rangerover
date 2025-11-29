package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

// Config holds configuration for the Downloader.
type Config struct {
	Url        string
	Dest       string
	Workers    int
	ChunkSize  int64
	MaxRetries int
	Timeout    time.Duration
}

// Validate checks if the configuration is valid.
func (c *Config) Validate() error {
	if c.Url == "" {
		return errors.New("url is required")
	}
	if c.Dest == "" {
		return errors.New("destination is required")
	}
	if c.Workers < 1 {
		return fmt.Errorf("workers must be at least 1, got %d", c.Workers)
	}
	if c.ChunkSize < 1 {
		return fmt.Errorf("chunk size must be positive, got %d", c.ChunkSize)
	}
	if c.MaxRetries < 0 {
		return fmt.Errorf("max retries cannot be negative, got %d", c.MaxRetries)
	}
	return nil
}

// SetDefaults sets default values for unspecified config fields.
func (c *Config) SetDefaults() {
	if c.Workers == 0 {
		c.Workers = 4
	}
	if c.ChunkSize == 0 {
		c.ChunkSize = 1 << 20 // 1MB
	}
	if c.MaxRetries == 0 {
		c.MaxRetries = 3
	}
	if c.Timeout == 0 {
		c.Timeout = 30 * time.Second
	}
}

type Downloader struct {
	cfg    Config
	Client *http.Client
}

// NewDownloader creates a new Downloader with the provided configuration.
func NewDownloader(cfg Config) (*Downloader, error) {
	// Set defaults for zero values
	cfg.SetDefaults()

	// Validate configuration
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	// Create HTTP client with reasonable defaults
	client := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        cfg.Workers,
			MaxIdleConnsPerHost: cfg.Workers,
			IdleConnTimeout:     90 * time.Second,
		},
	}

	return &Downloader{
		cfg:    cfg,
		Client: client,
	}, nil
}

// RangeHeaders gets Accept-Ranges, Content-Length, and Etag header using a Head http request
func RangeHeaders(url string) (RangesHeader, error) {
	resp, err := http.Head(url)
	if err != nil {
		return RangesHeader{}, fmt.Errorf("error heading url %s: %w", url, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return RangesHeader{}, fmt.Errorf("Status code is not okay it is %d\n", resp.StatusCode)
	}
	rangesHeaders := RangesHeader{
		resp.Header.Get("Accept-Ranges"),
		resp.ContentLength,
		resp.Header.Get("ETag"),
	}
	return rangesHeaders, nil
}

func (d *Downloader) Start() error {
	rangesHeader, err := RangeHeaders(d.cfg.Url)
	if err != nil {
		return err
	}

	fileSize := rangesHeader.ContentLength

	file, err := createEmptyFile(d.cfg.Dest, fileSize)
	if err != nil {
		return err
	}
	defer file.Close()

	// func calculateChunks(fileSize, ChunkSize) []Chunk

	chunk := Chunk{
		size:   min(fileSize, 1_048_576),
		offset: 0,
	}

	err = Download(d.cfg.Url, chunk, file)
	if err != nil {
		return err
	}

	return nil
}
