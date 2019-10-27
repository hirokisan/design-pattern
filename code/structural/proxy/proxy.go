package proxy

type Instance interface {
	MethodA() string
	MethodB() string
}

type InstanceA struct {
}

func NewInstanceA() Instance {
	return &InstanceA{}
}

func (i *InstanceA) MethodA() string {
	return "MethodA"
}

func (i *InstanceA) MethodB() string {
	return "MethodB"
}

type InstanceB struct {
	InstanceA Instance
}

func NewInstanceB() Instance {
	return &InstanceB{
		InstanceA: NewInstanceA(),
	}
}

func (i *InstanceB) MethodA() string {
	return "MethodA"
}

func (i *InstanceB) MethodB() string {
	return i.InstanceA.MethodB()
}
