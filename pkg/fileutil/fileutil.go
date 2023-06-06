package fileutil

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"time"
)

// CompareFiles compares two files and returns true if they have the same content.
func CompareFiles(file1, file2 string) (bool, error) {
	hash1, err := getFileHash(file1)
	if err != nil {
		return false, err
	}

	hash2, err := getFileHash(file2)
	if err != nil {
		return false, err
	}

	return hash1 == hash2, nil
}

// GetLastModified returns the last modified time of a file.
func GetLastModified(file string) (time.Time, error) {
	fi, err := os.Stat(file)
	if err != nil {
		return time.Time{}, err
	}

	return fi.ModTime(), nil
}

// getFileHash calculates the MD5 hash of a file.
func getFileHash(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
