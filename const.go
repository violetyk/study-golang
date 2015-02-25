// +build ignore

package main

import (
	"github.com/k0kubun/pp"
)

const (
	OK1 = "ok"
	OK2 = 100
	OK3 = true
	// NG  = []string{"aa", "bb", "cc", "dd"}
)

const (
	JPEG int = iota
	PNG
	GIF
	TIFF
	BMP
)

func main() {

	pp.Println(OK1, OK2, OK3)

	pp.Println(JPEG, PNG, GIF, TIFF, BMP)
	// pp.Println(NG)
}
