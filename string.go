// +build ignore

package main

import (
	"strconv"
	"strings"

	"github.com/k0kubun/pp"
)

func main() {
	var s string
	var i int

	s = "100"
	i, _ = strconv.Atoi(s)
	pp.Println(i)

	s = "1,000"
	i, _ = strconv.Atoi(s)
	pp.Println(i)

	s = "a"
	i, _ = strconv.Atoi(s)
	pp.Println(i)

	s = "1,2,3,4,"
	for _, v := range strings.Split(s, ",") {
		pp.Println(v)
	}

}
