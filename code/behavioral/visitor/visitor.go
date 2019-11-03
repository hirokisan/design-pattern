package visitor

import (
	"errors"
)

// Element : データ側
// visitorによって処理される対象
type Element interface {
	Accept(Visitor) error
}

type element struct {
	enabled bool
}

func NewElement() *element {
	return &element{}
}

func (e *element) Accept(visitor Visitor) error {
	return visitor.visit(e)
}

// Visitor : 訪問者
// elementに対する訪問者の要件を規定している
type Visitor interface {
	visit(*element) error
}

var (
	ErrFailed = errors.New("failed")
)

type visitorA struct {
}

func NewVisitorA() *visitorA {
	return &visitorA{}
}

func (v *visitorA) visit(element *element) error {
	if !element.enabled {
		return ErrFailed
	}

	return nil
}

type visitorB struct {
}

func NewVisitorB() *visitorB {
	return &visitorB{}
}

func (v *visitorB) visit(element *element) error {
	if element.enabled {
		return ErrFailed
	}

	return nil
}
