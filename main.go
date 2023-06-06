package main

import (
	"fmt"
	"math"
)

func main() {
	file1 := "/Users/marekgaj/Desktop/folder_1/some_empty_folder/another.txt"
	file2 := "/Users/marekgaj/Desktop/folder_2/some_empty_folder/another.txt"

	FileComp(file1, file2)
}

func FileComp(file1, file2 string) {
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

	fmt.Printf("Time difference between editions: %.0f minutes and %.0f seconds.\n", math.Floor(timeDiff.Minutes()), math.Floor(timeDiff.Seconds()))
}
