// +build ignore

package main

import (
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1) // goroutine を生成するたびインクリメント
		go func(i int) {
			log.Println(i)
			wg.Done() // 終了時にデクリメント
		}(i)
	}
	log.Println("waiting")
	wg.Wait() // ブロックし、全ての Done が終わったら次に進む
	log.Println("finish")
}
