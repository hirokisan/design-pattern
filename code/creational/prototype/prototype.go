package prototype

// TODO : interface にしてみる

type Prototype struct {
}

func (p *Prototype) Clone() *Prototype {
	clone := *p

	return &clone
}

func NewPrototype() *Prototype {
	return &Prototype{}
}
