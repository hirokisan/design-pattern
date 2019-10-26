package bridge_test

import (
	"testing"

	bridge "github.com/hirokisan/design-pattern/code/structural/bridge"
	"github.com/stretchr/testify/assert"
)

func TestBridge(t *testing.T) {
	obj := bridge.NewBridge(
		bridge.NewMethodImplA(),
		bridge.NewMethodImplB(),
	)

	assert.Equal(t, "this is MethodA", obj.RunMethodA())
	assert.Equal(t, "this is MethodB", obj.RunMethodB())
}
