package main

import "fmt"
import "reflect"

// structの定義
type MyKlass struct {
	Prop string
}

func (k MyKlass) Meth() {
	fmt.Printf("This prop is %s", k.Prop)
}

// 型情報を貯めておく
// キー文字列バリュー型情報のmap
var typeRegistry = map[string]reflect.Type{}

// 型情報を貯めるメソッド
func registerType(typ reflect.Type) {
	typeRegistry[typ.String()] = typ
}

// まずmainより前に
// MyKlassという型を登録しとく
func init() {
	registerType(reflect.TypeOf(MyKlass{}))
}
func main() {
	typeName := "main.MyKlass"

	// ちょっと複雑。
	// reflect.Type型を使ってNewして reflect.InterfaceValueをつくる
	// InterfaceValueをElemでValueにする
	// Interfaceメソッドでこれをinterface{}にする
	// interface{}がMyKlassの実装を満たしているか.(type)で調べる
	k, _ := reflect.New(typeRegistry[typeName]).Elem().Interface().(MyKlass)

	k.Prop = "foo"
	k.Meth()
}
