// +build ignore

package main

import (
	"log"
	"runtime"
)

func main() {
	// マルチコアで実行する
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)
	log.Println(cpus)
}
