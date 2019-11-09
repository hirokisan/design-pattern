package marker_interface_test

import (
	"testing"

	"github.com/hirokisan/design-pattern/other/marker_interface"
	"github.com/stretchr/testify/assert"
)

func TestMarkerInterface(t *testing.T) {
	assert.NoError(t, marker_interface.Accept("string"))
	assert.NoError(t, marker_interface.Accept(64))
	assert.Error(t, marker_interface.Accept([]int{64, 64}))
}
