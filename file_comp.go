package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"math"
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

func File_Comp() {
	file1 := "/Users/marekgaj/Desktop/folder_1/some_empty_folder/another.txt"
	file2 := "/Users/marekgaj/Desktop/folder_2/some_empty_folder/another.txt"

	sameContent, err := CompareFiles(file1, file2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if sameContent {
		fmt.Println("The files have the same content.")
	} else {
		fmt.Println("The files have different content.")
	}

	// Get last modified timestamps
	modTime1, err := GetLastModified(file1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	modTime2, err := GetLastModified(file2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Calculate the time difference
	timeDiff := modTime1.Sub(modTime2)

	fmt.Print("Time difference between edition:", math.Floor(timeDiff.Abs().Minutes()), " minutes and ", math.Floor(math.Floor(timeDiff.Abs().Minutes())-math.Floor(timeDiff.Abs().Minutes())), " seconds.")
}
