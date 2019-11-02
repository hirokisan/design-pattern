package memento_test

import (
	"testing"

	"github.com/hirokisan/design-pattern/code/behavioral/memento"
	"github.com/stretchr/testify/assert"
)

func TestMemento(t *testing.T) {
	obj := memento.NewOriginator()

	assert.Equal(t, 0, obj.FieldA)

	memento := obj.CreateMemento()

	obj.FieldA = 10

	assert.Equal(t, 10, obj.FieldA)

	obj.RestoreMemento(memento)

	assert.Equal(t, 0, obj.FieldA)
}
