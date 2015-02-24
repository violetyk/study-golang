// +build ignore

package main

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"io"
	"io/ioutil"
	"log"
	"os"
	"reflect"

	"github.com/gographics/imagick/imagick"
)

func main() {

	imagick.Initialize()
	defer imagick.Terminate()

	var wand *imagick.MagickWand = imagick.NewMagickWand()
	defer wand.Destroy()

	// orig_path := "/Users/kagaya/Desktop/takao1.jpg"
	orig_path := "/Users/kagaya/Desktop/test.jpg"
	dist_path := "/Users/kagaya/Desktop/takao1_resized_by_imagick.jpg"

	file, err := os.Open(orig_path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	i, err := jpeg.DecodeConfig(file)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	wand.SetImageFormat("jpeg")
	wand.SetOption("jpeg:size", fmt.Sprintf("%dx%d", i.Width, i.Height))
	log.Println(reflect.TypeOf(data))
	// err = wand.ReadImageBlob([]byte(data))
	err = wand.ReadImage(orig_path)
	if err != nil {
		panic(err)
	}

	// width := wand.GetImageWidth()
	// height := wand.GetImageHeight()
	hWidth := uint(i.Width / 2)
	hHeight := uint(i.Height / 2)
	log.Println(hHeight, hWidth)

	// Resize the image using the Lanczos filter
	// The blur factor is a float, where > 1 is blurry, < 1 is sharp
	err2 := wand.ResizeImage(hWidth, hHeight, imagick.FILTER_LANCZOS, 1)
	if err2 != nil {
		panic(err2)
	}

	// Set the compression quality to 95 (high quality = low compression)
	err = wand.SetImageCompressionQuality(95)

	// Exif情報を削除
	wand.StripImage()

	resized_data := wand.GetImageBlob()
	file2, err3 := os.Create(dist_path)
	if err3 != nil {
		panic(err3)
	}
	io.Copy(file2, bytes.NewReader(resized_data))

	wand.Clear()
}
