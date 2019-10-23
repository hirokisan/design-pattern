package main

import "fmt"

// Factory : インスタンス生成の大まかな処理
type Factory struct {
	CreateOutput func() Output
}

// Create : CreateOutput を用いてOutputを生成することが定義されているが、細かい作り方は定義していない
// NOTICE : これがtemplate methodパターンであり、template methodパターンを用いてfactoryを実現している
func (f *Factory) Create() Output {
	return f.CreateOutput()
}

// NewFactory : CreateOutputは外部から注入するなどできる
func NewFactory(f func() Output) *Factory {
	return &Factory{
		CreateOutput: f,
	}
}

type OutputImpl struct{}

func (o *OutputImpl) Method() {
	fmt.Println("implemented")
}

func CreateOutputImpl() Output {
	return &OutputImpl{}
}

// Output : 生成されるインスタンス
type Output interface {
	Method()
}

func main() {
	obj := NewFactory(CreateOutputImpl).Create()
	obj.Method()
}
