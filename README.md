study-golang
============

# Hello world

macにインストール

```
brew install go
go version
```

実行

```
go run hello.go
```

ビルド

```
go build hello.go
./hello
```

規約に合わせてソースを整形

```
go fmt hello.go
```

[クロスコンパイルのやり方](http://qiita.com/Jxck_/items/02185f51162e92759ebe)


# パスとディレクトリ
GOPATHとディレクトリ名の規約による依存関係解決

- bin/
- pkg/ 依存パッケージのオブジェクトファイル
- src/

ビルドして$GOPATH/bin/へインストール

```
go install
```


外部パッケージの取得

```
go get github.com/wdpress/gosample
```


GOPATHをマシンで1つに固定して開発するスタイルが一般的

```
export GOPATH=$HOME
export PATH=$PATH:$GOPATH/bin
```

# 基本文法
- mainパッケージ main()関数
- import、オプション
- 組み込み型
  - uint8, uint16, uint32, uint64 / 符号無し整数
  - uint / 32か4ビットの符号無し整数
  - int8, int16, int32, int64 / 符号あり整数
  - int / 32か4ビットの符号あり整数
  - float32, float64 / 浮動小数点
  - complex64, complex128 / 複素数
  - byte / uint8のエイリアス
  - rune / Unicodeのコードポイント、charのような1文字を表すことに使う。シングルクォートで囲む。
  - uintptr / ポインタ値用符号無し整数
  - error / エラーを表すインターフェース
- 文字列
  - 文字列はダブルクォートで囲う
  - ヒアドキュメントはバッククォートで囲う
- 変数
  - var 変数名 型
  - 複数の初期化と宣言
  - := / 関数内部で var と型宣言を省略できる記法
- 定数
  - var 定数名 型
  - 定数宣言できるのはerror型以外
- ゼロ値（初期化しなかったときのデフォルト値）
  - 整数 / 0 
  - 浮動小数点 / 0.0
  - bool / false
  - string / ""
  - 配列 / 各要素がゼロ値になってる配列
  - 構造体 / 各フィールドがゼロ値になってる構造
  - その他 / nil
- if / else if / else
  - 条件部の丸括弧無し 
  - 1行で省略するif文はなし
  - 三項演算子なし
- for
  - 繰り返しはすべてfor
  - 条件部の丸括弧無し 
  - for i := 1; i < 10; i++ { }
  - for n < 10 { } // whileの代わり
  - for { doSomething() } // 無限ループ
- swith
  - breakいらず
  - fallthrough で次のcaseに処理を移す
  - caseに式がかける
- 関数
  - func sum(i, j int) int {
  - 引数の型の宣言は最後に1つにまとめられる
  - 戻り値がある場合、引数の次に型を指定する
  - func swap(i, j int) (int, int) {
  - 複数の値をreturnで返せる
  - 受け取る変数の数が合わないとエラーだが、アンスコで明示的に無視できる
  - エラーを返す関数で、正常だった場合にはnilが返り、以上だった場合にはerror型の変数にだけ値が入る
  - 自作のエラーはerrorsパッケージを使って作る
  - return 0, errors.New("div by zero")
    - エラーを返す関数で、正常だった場合にはnilが返り、以上だった場合にはerror型の変数にだけ値が入る
    - 自作のエラーはerrorsパッケージを使って作る
  - 名前付き戻り値
  - func div(i, j int) (return r, err error) {
    - 名前付き戻り値とゼロ値初期化を使って、returnの書き間違えを防ぐテクニック
  - 可変長引数
    - func sum_all(nums ...int) (result int) { //numsは[]int型
- 無名関数
  - 関数リテラルを使って無名関数を作成
  - func(引数情報) { 処理 } (引数)
  - 関数を変数に代入したり引数に渡すことができる（Goの関数は第一級オブジェクト）
- 配列
  - 固定長
  - 要素番号はゼロから
  - var 変数名 [長さ]型
  - var arr [4]string
  - arr2 := [4]string{"a", "b", "c", "d"}
  - arr3 := [...]string{"a", "b", "c", "d"}
  - 関数に配列を渡す
    - 配列の型は長さを含むので、長さが違う配列を引数に渡そうとするとコンパイルエラー
    - コピー渡しになる
- スライス
  - 要は可変長配列
  - var s1 []string
  - s2 := []string{"aa", "bb", "cc", "dd"}
  - append()で末尾に値を追加
    - s1 = append(s1, "aa", "bb")
- 添え字アクセス
  - s2[0]
  - s2[2:3]
  - s2[2:]
  - s2[:3]
  - s2[1:len(s2)]
  - s2[:]
- range
  - 先頭から走査するために使う
  - for i, s := range s2 {
  - 配列やスライスの他、stringやマップ、チャネルにも使用可能
- マップ
  - 要は連想配列
  - var month map[int]string = map[int]string{}
  - month := map[int]string{ 1: "January", 2: "February" }
  - 2番目の値で存在確認
    - `_, ok := month[3]`
  - keyの削除はdelete()で
  - マップを走査するにはforとrangeで
    - for key, value := range month {
    - 順番は保証されないので注意
- ポインタ
  - 型の前に`*`を付ける
  - アドレスは変数の前に`&`を付ける
- defer
  - その関数を抜ける前に必ず実行される処理を記述できる
- パニック
  - 例外のようなもの
  - 組み込み関数 recover()でエラーを取得できる
  - defer 関数 とかしておいてパニックを補足し、その関数の中でrecover()でエラーを取得する
  - パニックが起きる前に補足する処理を書いておく必要があった
  - panic()を使って自前で発生することもできるが、基本的にはエラーを関数の戻り値として返すのがよい


# type
  - type MyPoint int
  - var p MyPoint = 100
  - func Add(p MyPoint, points)
  - 既存の型を拡張して独自の型を定義
  - 適切な型を用意して型レベルでコンパイル時にチェックすることで、堅牢なプログラムにできる


# 構造体(struct)
- メソッドが定義できるのでクラスに近い
- type 構造体名 struct { フィールド名 型 }
- フィールド
  - フィールド名が大文字で始まる場合はパブリック
  - フィールド名が小文字で始まる場合は閉じたスコープ
  - 初期化時にフィールドに値を渡さなかったらその型のゼロ値がセットされる
  - フィールドにはドット(.)でアクセス
- 構造体のポインタ型
  - 関数に構造体をコピー渡ししたくないときに使ったりする
  - `var task4 *Task = &Task{4, "buy the apple", false}`
- new()
  - 構造体の全フィールドをゼロ値で初期化して、そのポインタを返す
    - `var task4 *Task = new(Task)`
- コンストラクタ
  - Goにはない
  - Newで始まる関数を定義して構造体のポインタ型を返すのが通例
  - `f.Printf("%+v", task5)` で構造体の文字列表現を出力
  - String()メソッドを定義すると構造体文字列表現の上書きができるっぽい
- メソッド
  - メソッドを実行した対象の型をレシーバとして受け取り、func内で使用できる
  - レシーバをポインタにすれば、呼び出し側の構造体に変更を反映できる


# インタフェース
- その型がどのようなメソッドを実装すべきかを規定する役割
- implements みたいな構文はなく、型に定義されたメソッドがあるかどうかで実装しているかをみている


# 構造体、インタフェースの埋込み(Embed)
- Goには継承がない
- 構造体に他の型を埋め込んで、構造体やインタフェースの振る舞いを拡張できる
- 埋め込まれた型のフィールドやメソッド、型の実体にアクセスできる
- インタフェースの埋め込みは、複数のインタフェースから新しいインタフェースを定義する為に使ったりする
  - ReaderとWriterを埋め込んだReaderWriter
  - ReaderとCloserを埋め込んだReadCloser

# キャスト
- キャスト先の型(値)
- キャストに失敗するとpanic

# Type Assertion、Type Switch
- 型を調べる
  - s, ok := value.(string)
  - 第一戻り値には調べた型に変換された値
  - 第二戻り値には判定結果
- switchと組み合わせて型による分岐処理ができる（Type Switch）
  - switch v := value.(type) {

# 標準パッケージ
- encoding/json パッケージ
  - json.Marshal(構造体)
  - 小文字で始まるプライベートなフィールドは、基本的にJSONに含まれない
  - タグ
    - ``json:"name"`       // nameというキーで格納`
    - ``json:"-"`          // JSONに格納しない`
    - ``json:",omitempty"` // 値が空なら無視`
    - ``json:",string"`    // 値をJSONとして格納`
    - ダブルクォーテーションが抜けてもコンパイルエラーにならず、タグで指定した処理も動かなかったので注意
  - err := json.Unmarshal(JSON文字列, &格納先構造体のポインタ)
- osパッケージ ファイルの作成やオープン
  - `file, err := os.Create("ファイルのパス") // os.File構造体へのポインタを返す`
  - `defer file.Close() // 関数の終わりでdeferによって必ずファイルを閉じている`
  - `*os.File`はio.ReadWriteCloserインタフェースを実装している
  - `file, err := os.Open("すでにあるファイルのパス")`
- ioパッケージ ファイルの読み書き
  - Write([]byte), WriteString(string)
  - file.Read([]byte)
  - `content := make([]byte, 20) // 20バイトのスライスを作成`
  - JSONのEncoder/Decoderは引数にio.ReaderWriterを扱うAPIとなっている
  - そのメソッドにio.ReadWriterである`*os.File`を渡してあげればファイルへJSONの読み書きができる
  - 標準ライブラリには、io.Reader, io.Writerを中心として設計したAPIが多くある
- io/ioutilパッケージは簡単ファイル操作ができて便利
  - `ioutil.ReadAll(*os.File)`
  - `ioutil.ReadFile("./file.txt")`
  - `ioutil.WriteFile("./file.txt", bs2, 0666)`
- net/httpパッケージ HTTPサーバやクライアントを作る
  - `http.HandleFunc()`でルーティング
  - `http.IndeHandler(w http.ResponseWriter, r *http.Request)`
    - 第二引数にはリクエスト情報が入っている
    - 組み立てた結果をResponseWriterに書き込む
  - `http.ListenAndServe("3000", nil)`
