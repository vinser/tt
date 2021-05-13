package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	scanDir("/books")

}

// Scan dir
func scanDir(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	entries, err := f.Readdir(-1)
	if err := f.Close(); err != nil {
		return err
	}
	for _, entry := range entries {
		switch {
		case entry.IsDir():
			fmt.Println("===", path+string(os.PathSeparator)+entry.Name())
			scanDir(path + string(os.PathSeparator) + entry.Name())
		case strings.ToLower(filepath.Ext(entry.Name())) == ".zip":
			fmt.Println("Zip: ", entry.Name())
			processZip(path + string(os.PathSeparator) + entry.Name())
		case strings.ToLower(filepath.Ext(entry.Name())) == ".fb2":
			fmt.Println("FB2: ", entry.Name())
		}
	}
	return nil
}

// Get zip
func processZip(zipPath string) {
	zr, err := zip.OpenReader(zipPath)
	if err != nil {
		panic(err)
	}

	for _, file := range zr.File {
		fmt.Println("===========================================")
		fmt.Println("File               : ", file.Name)
		fmt.Println("NonUTF8            : ", file.NonUTF8)
		fmt.Println("Modified           : ", file.Modified)
		fmt.Println("CRC32              : ", file.CRC32)
		fmt.Println("UncompressedSize64 : ", file.UncompressedSize)
		fmt.Println("===========================================")
		rc, _ := file.Open()
		processFB2(rc)
	}
}

// Read file
func processFB2(io.ReadCloser) {

}

//
