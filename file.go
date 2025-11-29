package main

import (
	"io"
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

	if _, err := file.Seek(size-1, io.SeekStart); err != nil {
		file.Close()
		return nil, err
	}
	if _, err := file.Write([]byte{0}); err != nil {
		file.Close()
		return nil, err
	}
	return file, nil
}
