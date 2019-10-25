package prototype_test

import (
	"testing"

	"github.com/hirokisan/design-pattern/code/creational/prototype"
	"github.com/stretchr/testify/assert"
)

func TestPrototype(t *testing.T) {
	obj := prototype.NewPrototype()
	clone := obj.Clone()

	assert.Exactly(t, *obj, *clone, "等価である")
	assert.True(t, obj != clone, "等値ではない")
}

func TestPrototypeWithModify(t *testing.T) {
	obj := prototype.NewPrototype()
	obj.Name("me")
	clone := obj.Clone()
	clone.Name("you")

	assert.NotEqual(t, *obj, *clone, "等価ではない")
	assert.True(t, obj != clone, "等値ではない")
}
