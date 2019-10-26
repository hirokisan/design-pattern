package composite

import "fmt"

// Component : 再帰的な構造パターンをinterfaceとして抽出する
type Component interface {
	Greet()
}

type Person struct {
	Name string
}

func NewPerson(name string) *Person {
	return &Person{
		Name: name,
	}
}

func (p *Person) Greet() {
	fmt.Println(p.Name)
}

type Composite struct {
	List []Component
}

func NewComposite() *Composite {
	return &Composite{
		List: make([]Component, 0),
	}
}

func (c *Composite) Add(component Component) {
	c.List = append(c.List, component)
}

func (c *Composite) Greet() {
	for _, component := range c.List {
		component.Greet()
	}
}
