// +build ignore

package main

import (
	"fmt"
	"runtime"
	"time"
)

func finalizer(m *string) {
	fmt.Println("finalize for ->", *m)
	return
}

func createInstance() {
	x := "hello"
	runtime.SetFinalizer(&x, finalizer)
	return
}

func main() {
	createInstance()
	runtime.GC()
	time.Sleep(1 * time.Second)
	return
}
