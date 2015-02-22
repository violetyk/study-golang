// +build ignore

package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func main() {

	// 現在起動しているgoroutineの数

	// main自体が1個のgoroutine
	// あとruntime.MHeap_Scavenger() が動いてる
	fmt.Println(runtime.NumGoroutine())

	go func() {
		log.Println("end")
		fmt.Println(runtime.NumGoroutine())
	}()

	go func() {
		log.Println("return")
		fmt.Println(runtime.NumGoroutine())
		return
	}()

	go func() {
		log.Println("exit")
		fmt.Println(runtime.NumGoroutine())
		runtime.Goexit()
	}()

	time.Sleep(time.Second)

	// log.Println(runtime.NumGoroutine())
	// buf := make([]byte, 1<<20)
	// buf = buf[:runtime.Stack(buf, true)]
	// log.Println(string(buf))
	// select {}
}
