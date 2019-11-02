package memento

// Memento : 状態を保持する
type memento struct {
	fieldA int
}

type Originator interface {
	CreateMemento() *memento
	RestoreMemento(memento)
}

type originator struct {
	FieldA int
}

func NewOriginator() *originator {
	return &originator{}
}

func (o *originator) CreateMemento() *memento {
	obj := &memento{
		fieldA: o.FieldA,
	}

	return obj
}

func (o *originator) RestoreMemento(memento *memento) {
	o.FieldA = memento.fieldA
}
