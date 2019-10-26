package decorator

type Interface interface {
	Method() string
}

type Instance struct {
}

func NewInstance() *Instance {
	return &Instance{}
}

func (i *Instance) Method() string {
	return "this is base method"
}

type Decorator struct {
	Instance Interface
}

func NewDecorator(instance Interface) Interface {
	return &Decorator{
		Instance: instance,
	}
}

func (d *Decorator) Method() string {
	return d.Instance.Method() + " from decorator"
}
