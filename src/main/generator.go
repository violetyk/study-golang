// generatorは扱う値を実行時に作るので省メモリ
package main

import "log"

func generator(n int) chan int {
	// 同時に1個までしか作成されないことを保証する
	ch := make(chan int)
	i := 0
	go func() {
		for {
			ch <- i
			i++
			if i > n {
				close(ch)
				break
			}
		}
	}()
	return ch
}

func main() {
	for x := range generator(10) {
		log.Println(x)
	}
}
