package filetransfer

import (
	"fmt"
	"io"
	"os"
)

// TransferFile transfers the content of one file to another.
func TransferFile(sourceFile, destinationFile string) error {
	source, err := os.Open(sourceFile)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(destinationFile)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err != nil {
		return err
	}

	return nil
}

// ResolveConflict resolves conflicts between the source and destination files.
func ResolveConflict(sourceFile, destinationFile string) error {
	sourceInfo, err := os.Stat(sourceFile)
	if err != nil {
		return err
	}

	destinationInfo, err := os.Stat(destinationFile)
	if err != nil {
		return err
	}

	sourceModifiedTime := sourceInfo.ModTime()
	destinationModifiedTime := destinationInfo.ModTime()

	// Compare the modified times of the files
	if sourceModifiedTime.After(destinationModifiedTime) {
		// Source file is more recent, overwrite the destination file
		err = TransferFile(sourceFile, destinationFile)
		if err != nil {
			return err
		}
		fmt.Println("Conflict resolved: Destination file overwritten with the source file.")
	} else if sourceModifiedTime.Before(destinationModifiedTime) {
		// Destination file is more recent, keep the destination file
		fmt.Println("Conflict resolved: Source file ignored, destination file remains unchanged.")
	} else {
		// Files have the same modified time, no conflict
		fmt.Println("No conflict: Both files have the same modified time.")
	}

	return nil
}
