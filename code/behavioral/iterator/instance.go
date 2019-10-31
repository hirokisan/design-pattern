package iterator

type Instance interface {
	Method() bool
}

type Inst struct {
}

func NewInst() *Inst {
	return &Inst{}
}

func (i *Inst) Method() bool {
	return true
}
