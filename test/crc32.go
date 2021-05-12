package main

import (
	"fmt"
	"hash/crc32"
	"io/ioutil"
)

func FileCRC32(filePath string) (uint32, error) {
	fbytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return 0, err
	}
	return crc32.ChecksumIEEE(fbytes), nil
}

func main() {

	fmt.Println(FileCRC32("/books/577876.fb2"))
}
