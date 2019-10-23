package main

import "fmt"

// Template : 大まかな処理は定義しているが、具体的な処理の内容は定義していない
type Template struct {
	Method func()
}

func (t *Template) Run() {
	t.Method()
}

func NewObject(f func()) *Template {
	return &Template{
		Method: f,
	}
}

// Impl : Methodを実装する
func Impl() {
	fmt.Println("implemented")
}

func main() {
	obj := NewObject(Impl)
	obj.Run()
}
