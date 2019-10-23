package main

import "fmt"

// Interface : 大まかな処理は定義しているが、具体的な処理の内容は定義していない
type Interface interface {
	Method()
}

// Impl : Interface の処理を実装する
type Impl struct {
}

func NewImpl() Interface {
	return &Impl{}
}

func (i *Impl) Method() {
	fmt.Println("Interface is implemented by Impl")
}

func main() {
	obj := NewImpl()
	obj.Method()
}
