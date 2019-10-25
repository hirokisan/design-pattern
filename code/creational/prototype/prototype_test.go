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
