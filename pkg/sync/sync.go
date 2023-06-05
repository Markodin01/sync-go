package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	sourceDir := "/Users/marekgaj/Desktop/folder_1"
	destinationDir := "/Users/marekgaj/Desktop/folder_2"

	File_Comp()
	err := SyncDirectories(sourceDir, destinationDir)
	if err != nil {
		fmt.Println("Error:", err)
	}
	File_Comp()
}

func SyncDirectories(sourceDir, destinationDir string) error {
	return filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relativePath, err := filepath.Rel(sourceDir, path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(destinationDir, relativePath)

		if info.IsDir() {
			// Create the corresponding directory in the destination if it doesn't exist
			err := os.MkdirAll(destPath, info.Mode())
			if err != nil {
				return err
			}
		} else {
			// Copy the file from source to destination
			err := CopyFile(path, destPath)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func CopyFile(source, destination string) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}
