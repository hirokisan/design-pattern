package proxy_test

import (
	"testing"

	"github.com/hirokisan/design-pattern/code/structural/proxy"
	"github.com/stretchr/testify/assert"
)

func TestProxy(t *testing.T) {
	obj := proxy.NewInstanceB()

	assert.Equal(t, "MethodB", obj.MethodB())
}
