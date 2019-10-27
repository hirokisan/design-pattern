package facade_test

import (
	"testing"

	"github.com/hirokisan/design-pattern/code/structural/facade"
	"github.com/stretchr/testify/assert"
)

func TestFacade(t *testing.T) {
	obj := facade.NewFacade()

	assert.Equal(t, "MethodA and MethodB called", obj.Method())
}
