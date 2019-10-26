package decorator_test

import (
	"testing"

	"github.com/hirokisan/design-pattern/code/structural/decorator"
	"github.com/stretchr/testify/assert"
)

func TestDecorator(t *testing.T) {
	obj := decorator.NewInstance()
	decorator := decorator.NewDecorator(obj)

	assert.Equal(t, "this is base method", obj.Method())
	assert.Equal(t, "this is base method from decorator", decorator.Method())
}
