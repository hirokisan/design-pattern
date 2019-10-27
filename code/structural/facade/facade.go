package facade

import "fmt"

// Facade : クライアント視点で使いやすいAPIを提供する
type Facade struct {
}

func NewFacade() *Facade {
	return &Facade{}
}

// Method : 複数の処理をまとめる
func (f *Facade) Method() string {
	return fmt.Sprintf("%s and %s called", MethodA(), MethodB())
}

func MethodA() string {
	return "MethodA"
}

func MethodB() string {
	return "MethodB"
}
