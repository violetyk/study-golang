package main

import (
  "fmt"
  "net/http"
)


func main() {

  // ルーティングの設定
  http.HandleFunc("/", IndexHandler)
  http.ListenAndServe(":3000", nil)


  fmt.Println("wef")
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Hello world")
}
