package main

import (
  // "fmt"
  "gosample"
  // "github.com/wdpress/gosample" // $GOPATH/src/からのパス
  f "fmt" // パッケージ名を変える
  _ "github.com/wdpress/gosample" // インポートしたパッケージを使わない
  . "strings" // パッケージ名を省略する
  _ "os"
  _ "log"
)

// 構造体。大文字から始まるフィールド名はパブリック
type Task struct {
  ID       int
  Detail   string
  done     bool
  *User // embedするときはフィールド名ではなく型のみを記述
}

// コンストラクタ代わりに、Newで始まる関数を定義して構造体のポインタ型を返す
func NewTask(id int, detail, firstName, lastName string) *Task {
  task := &Task {
    ID: id,
    Detail: detail,
    done: false,
    User: NewUser(firstName, lastName),
  }
  return task
}

// 構造体をポインタ型で受け取る
func Finish(task *Task) {
  task.done = true
}

// task.Finish()とできるようにレシーバを構造体のポインタとして受け取るようにする
// func (task *Task) Finish() {
//こんな感じでレシーバのポインタを返してあげるとメソッドチェーンできたけどあってるのか不安
func (task *Task) Finish() *Task {
  task.done = true
  return task
}

func (task *Task) IsFinish() bool {
  return task.done
}

// メソッドを実行した対象の型をレシーバとして受け取り、func内で使用できる
func (task Task) String() string {
  str := f.Sprintf("%d) %s", task.ID, task.Detail)
  return str
}

// インタフェースが単純な名前の場合、その関数にerを付ける慣習がある
type Stringer interface {
  String() string
}

// インタフェースを引数に取る
func Print(stringer Stringer) {
  f.Println(stringer.String())
}

// どんな型でも受け付ける関数を定義
type Any interface {
}
func Do(e Any) {
}
// func Do(e interface{}) { // インタフェースを定義しなくても直接かける
// }
// func Do(e ...interface{}) { // 任意の型の引数を可変長で受け取る
// }


// TaskにembedするUser情報
type User struct {
  FirstName string
  LastName  string
}

func (user *User) FullName() string {
  fullname := f.Sprintf("%s %s", user.FirstName, user.LastName)
  return fullname
}

func NewUser(firstName, lastName string) *User {
  return &User {
    FirstName: firstName,
    LastName: lastName,
  }
}



func main() {
  // fmt.Println(gosample.Message)
  f.Println(ToUpper(gosample.Message)) // strings.ToUpper

  message := "hello world"
  f.Println(message)


  // if
  a, b := 10, 100
  if a > b {
    f.Println("a > b")
  } else if a < b {
    f.Println("a < b")
  } else {
    f.Println("a = b")
  }


  // for
  for i := 1; i < 10; i++ {
    f.Println(i)
  }

  n := 0
  for n < 10 {
    f.Printf("n = %d\n", n)
    // ++n
    n++
  }


  // switch
  x := 10
  switch x {
  case 15:
    f.Println("FizzBuzz")
  case 5, 10:
    f.Println("Buzz")
  case 3, 6, 9:
    f.Println("Fizz")
  default:
    f.Println(n)
  }

  y := 3
  switch y {
  case 3:
    y--
    fallthrough
  case 2:
    y--
    fallthrough
  case 1:
    y--
    // fallthrough エラー
    f.Println(y)
  }

  z := 10
  switch {
  case z%15 == 0:
    f.Println("FizzBuzz")
  case z%5 == 0:
    f.Println("Buzz")
  case z%3 == 0:
    f.Println("Fizz")
  default:
    f.Println(z)
  }


  // 関数
  f.Println(sum(11,22))

  aa, bb := 3,4
  aa, bb = swap(aa, bb)
  f.Println(aa, bb)
  // aa = swap(aa, bb) エラー
  aa, _ = swap(aa, bb)
  f.Println(aa)


  // 無名関数
  func(i, j int) {
    f.Println(i + j)
  } (2, 4)


  // 変数に関数を代入
  // var 変数名 型
  var myfunc func(i, j int) = func(i, j int) {
    f.Println(i / j)
  }

  myfunc(1000, 10)

  // 配列
  // var 変数名 [サイズ]型
  var arr [4]string
  arr[0] = "a"
  arr[1] = "b"
  arr[2] = "c"
  arr[3] = "d"
  f.Println(arr[0])

  arr2 := [4]string{"a", "b", "c", "d"}
  f.Println(arr2[0])
  arr3 := [...]string{"a", "b", "c", "d"}
  f.Println(arr3[0])

  // スライス
  var s1 []string

  // s1[0] = "aa" エラー
  s1 = append(s1, "aa")
  s1 = append(s1, "bb")
  s1 = append(s1, "cc", "dd")
  f.Println(s1)

  s2 := []string{"aa", "bb", "cc", "dd"}
  f.Println(s2[0])
  f.Println(s2[2:3])
  f.Println(s2[2:])
  f.Println(s2[:3])
  f.Println(s2[1:len(s2)])
  f.Println(s2[:])

  // range
  for i, s := range s2 {
    f.Printf("%d = %s\n", i, s)
  }

  // 可変長引数
  f.Println(sum_all(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

  // map
  var month map[int]string = map[int]string{}
  month[1] = "January"
  month[2] = "February"
  f.Println(month)

  month2 := map[string]int{
    "January": 1,
    "February": 2,
  }
  f.Println(month2)

  _, ok := month[2]
  if ok {
    f.Println("exists!")
  }
  
  delete(month, 2)
  f.Println(month)

  for key, value := range month2 {
    f.Println(key, value)
  }

  // ポインタ
  var i = 10
  callByValue(i)
  f.Println(i)
  callByRef(&i) // アドレスを渡す
  f.Println(i)


  // defer, panic, recover()
  defer func () {
    err := recover()
    if err != nil {
      f.Println("パニックです")
      // log.Fatal(err)
    }
  }()
  // var s3 []string
  // f.Println(s3[0]) // panic

  // panic(errors.New("パニック発生"))

  // 構造体 
  var task Task = Task{
    ID: 1,
    Detail: "buy the milk",
    // done: false,
  }
  f.Println(task.ID)
  f.Println(task.Detail)
  f.Println(task.done)

  // var task2 Task = Task{2, "buy the banana"}
  // f.Println(task2.ID)
  // f.Println(task2.Detail)
  // f.Println(task2.done)

  // 明示的に指定しない場合、各フィールドゼロ値になる
  task3 := Task{}
  f.Println(task3.ID)
  f.Println(task3.Detail)
  f.Println(task3.done)

  // 構造体Taskのポインタ型
  // var task4 *Task = &Task{4, "buy the apple", false}
  var task4 *Task = new(Task)
  Finish(task4)
  f.Println(task4.done)
  
  // New+構造体名の関数でコンストラクタ的にするのが通例。構造体のポインタを返してあげる
  task5 := NewTask(5, "buy the coffee", "Taro", "Yamada")
  Finish(task5)
  f.Printf("%+v\n", task5)

  // 型に定義したメソッド
  f.Printf("%s\n", task)
  // task.Finish()
  // f.Printf("%+v\n", task)
  f.Println(task.Finish().IsFinish())

  // Task型に埋め込んだUser型へアクセス
  f.Println(task5.FirstName) // User型のフィールド
  f.Println(task5.LastName)
  f.Println(task5.FullName()) // User型のメソッド
  f.Println(task5.User) // User型自体

  
  // キャスト
  var ii uint8 = 3
  var jj uint32 = uint32(ii)
  f.Println(jj)

  var ss = "abc"
  var byby []byte = []byte(ss)
  f.Println(byby)


  IsString("hogefuga")
  IsString(123)
}


func IsString(value interface{}) {
  s, ok := value.(string) // Type Assertion 判定が成功したら、第一戻り値には変換された値が返る
  if ok {
    f.Printf("value is string: %s\n", s)
  } else {
    f.Printf("value is not string\n")
  }

  // Type Switch
  switch v := value.(type) {
  case string:
    f.Printf("value is string: %s\n", v)
  case int:
    f.Printf("value is int: %d\n", v)
  case Stringer:
    f.Printf("value is Stringer: %s\n", v)
  }
}


func sum(i, j int) int {
  return i + j
}

func swap(i, j int) (int, int) {
  return j, i
}

func sum_all(nums ...int) (result int) {
  for _, n := range nums {
    result += n
  }
  return
}

func callByValue(i int) {
  i = 20
}

func callByRef(i *int) {
  *i = 20 // 参照先を上書き
}
