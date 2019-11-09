package immutable

// Immutable : 不変
type Immutable struct {
	name string
}

func NewImmutable(name string) *Immutable {
	return &Immutable{
		name: name,
	}
}

func (i *Immutable) Name() string {
	return i.name
}
