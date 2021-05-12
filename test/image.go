package main

import (
	"encoding/base64"
	"fmt"
	"image/jpeg"
	"os"

	"github.com/nfnt/resize"
)

func main() {
	// fi, _ := os.Open("b64image.txt")
	fi, _ := os.Open("1.txt")
	defer fi.Close()
	// data, _ := io.ReadAll(base64.NewDecoder(base64.StdEncoding, fi))
	iImg, _ := jpeg.Decode(base64.NewDecoder(base64.StdEncoding, fi))
	b := iImg.Bounds()
	fmt.Print(b)
	oImg := resize.Resize(100, 0, iImg, resize.NearestNeighbor)
	fo, _ := os.OpenFile("4.jpg", os.O_WRONLY|os.O_CREATE, 0777)
	defer fo.Close()
	// fo.Write(data)
	jpeg.Encode(fo, oImg, nil)

}
