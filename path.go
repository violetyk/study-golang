// +build ignore

package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/k0kubun/pp"
)

var str string = "Test String Replace Replace"

func main() {

	// str := "Test String Replace Replace"

	// Replacerを使って置換する
	r := strings.NewReplacer("Test", "Tes", "String", "str", "Replace", "rep")
	resStr := r.Replace(str)
	pp.Println(resStr)

	s := "/var/tmp/graid/wefwef"
	pp.Println(filepath.Base(s))
	pp.Println(filepath.Dir(s))
	// pp.Println(utf8.DecodeRuneInString(os.PathSeparator))
	pp.Println(string(os.PathSeparator))

	// os.Chmod("/var/tmp/graid", 0777)
	pp.Println("finish")
}
