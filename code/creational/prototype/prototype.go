package prototype

// NOTE : 役割的にもinterfaceとの相性はよくなさそう

type Prototype struct {
	name string
}

func (p *Prototype) Clone() *Prototype {
	clone := *p

	return &clone
}

func (p *Prototype) Name(name string) {
	p.name = name
}

func NewPrototype() *Prototype {
	return &Prototype{}
}
