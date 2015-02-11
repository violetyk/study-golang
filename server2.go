package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"
)

func serve(w http.ResponseWriter, r *http.Request) {

	log.Println(runtime.NumGoroutine())
	fmt.Fprintln(w, "同時実行テストだよ")
	time.Sleep(5 * time.Second)
}

func main() {
	http.HandleFunc("/", serve)
	http.ListenAndServe(":1234", nil)
}
