package adapter

type Adapter interface {
	Method()
}

type Implement struct {
}

func NewImplement() *Implement {
	return &Implement{}
}

func (i *Implement) method() string {
	return "implemented"
}

// Method : 元々はmethodしか定義されていなかったのでmethodを使うMethodを用意した
func (i *Implement) Method() string {
	return i.method()
}
