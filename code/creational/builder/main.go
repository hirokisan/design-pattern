package main

import "fmt"

// Output : DirectorがBuilderを用いて生成した結果
type Output struct {
	ID int
}

// Builder : 材料（表現形式）を用意し、結果を生成する機能を持つ
type Builder interface {
	SetID(int) Builder
	Build() *Output
}

// Director : builderを使って結果を生成する役割を担う
type Director struct{}

func NewDirector() *Director {
	return &Director{}
}

func (d *Director) BuildOutput(builder Builder) *Output {
	return builder.SetID(3).Build()
}

// BuilderFirst : Builderインターフェイスの実装
type BuilderFirst struct {
	ID int
}

func NewBuilderFirst() Builder {
	return &BuilderFirst{}
}

func (b *BuilderFirst) SetID(id int) Builder {
	b.ID = id

	return b
}

func (b *BuilderFirst) Build() *Output {
	return &Output{
		ID: b.ID,
	}
}

func main() {
	builder := NewBuilderFirst()
	output := NewDirector().BuildOutput(builder)
	fmt.Println(output)
}
