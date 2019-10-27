package flyweight_test

import (
	"testing"

	"github.com/hirokisan/design-pattern/code/structural/flyweight"
	"github.com/stretchr/testify/assert"
)

func TestFlyweight(t *testing.T) {
	obj1 := flyweight.GetOrCreateInstance("key")
	obj2 := flyweight.GetOrCreateInstance("key")
	obj3 := flyweight.GetOrCreateInstance("keyOther")

	assert.Exactly(t, *obj1, *obj2)
	assert.True(t, obj1 == obj2)
	assert.Equal(t, "accessible", obj1.Access())

	assert.True(t, obj1 != obj3)
}
