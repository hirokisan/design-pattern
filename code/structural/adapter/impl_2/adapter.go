package adapter

type Adapter interface {
	Method()
}

type Material struct{}

func NewMaterial() *Material {
	return &Material{}
}

func (m *Material) method() string {
	return "implemented"
}

type Implement struct {
	Material *Material
}

func NewImplement() *Implement {
	return &Implement{
		Material: NewMaterial(),
	}
}

func (i *Implement) Method() string {
	return i.Material.method()
}
