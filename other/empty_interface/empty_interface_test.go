package empty_interface_test

import (
	"testing"

	"github.com/hirokisan/design-pattern/other/empty_interface"
	"github.com/stretchr/testify/assert"
)

func TestMarkerInterface(t *testing.T) {
	assert.NoError(t, empty_interface.Accept("string"))
	assert.NoError(t, empty_interface.Accept(64))
	assert.Error(t, empty_interface.Accept([]int{64, 64}))
}
