// +build ignore

package main

import (
	"image/jpeg"
	"io/ioutil" // "io/ioutil"
	"os"
)

func main() {

	orig_path := "/Users/kagaya/Desktop/takao1.jpg"
	dist_path := "/Users/kagaya/Desktop/takao1_resized_by_imaging.jpg"

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

}
