package main

import (
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  "os"
)

type Person struct {
  ID int        `json:"id"`
  Name string   `json:"name"`
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Hello world")
}

func PersonHandler(w http.ResponseWriter, r *http.Request) {
  defer r.Body.Close() // 処理の最後にBodyを閉じる
  
  if r.Method == "POST" {
    // リクエストボディをJSONに変換
    var person Person
    // decoder := json.NewDecoder(r.Body)
    // fmt.Println(r.Body)

    // err := decoder.Decode(&person)
    err := json.NewDecoder(r.Body).Decode(&person)
    if err != nil {
      log.Fatal(err)
    }

    // ファイル名を{id}.txtとする
    filename := fmt.Sprintf("%d.txt", person.ID)
    file, err := os.Create(filename)
    if err != nil {
      log.Fatal(err)
    }
    defer file.Close()

    // ファイルにNameを書き込む
    _, err = file.WriteString(person.Name)
    if err != nil {
      log.Fatal(err)
    }

    // レスポンスとしてステータスコード201を送信
    w.WriteHeader(http.StatusCreated)
  }

}

func main() {

  // ルーティングの設定
  http.HandleFunc("/", IndexHandler) // ルートをIndexHandler関数で受け取る
  http.HandleFunc("/persons", PersonHandler)
  http.ListenAndServe(":3000", nil)

  // POSTしてみる
  // curl localhost:3000/persons -d '{"id":1, "name":"gopher"}'


}

