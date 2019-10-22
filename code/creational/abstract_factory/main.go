package main

import "fmt"

type Factory interface {
	CreateInstance() Instance
}

type Instance interface {
	FirstMethod()
}

func NewAbstractFactory(name string) Factory {
	switch name {
	case "first":
		return NewFactoryFirst()
	case "second":
		return NewFactorySecond()
	default:
		panic(fmt.Sprintf("given name %s is not supported", name))
	}
}

type FactoryFirst struct{}

func NewFactoryFirst() *FactoryFirst {
	return &FactoryFirst{}
}

type InstanceFirst struct{}

func (f *FactoryFirst) CreateInstance() Instance {
	return &InstanceFirst{}
}

func (i *InstanceFirst) FirstMethod() {
	fmt.Println("this is firstMethod from InstanceFirst")
}

type FactorySecond struct{}

func NewFactorySecond() *FactorySecond {
	return &FactorySecond{}
}

type InstanceSecond struct{}

func (f *FactorySecond) CreateInstance() Instance {
	return &InstanceSecond{}
}

func (i *InstanceSecond) FirstMethod() {
	fmt.Println("this is firstMethod from InstanceSecond")
}

func main() {
	af := NewAbstractFactory("third")
	i := af.CreateInstance()
	i.FirstMethod()
}
