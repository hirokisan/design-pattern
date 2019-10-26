package adapter_test

import (
	"testing"

	adapter "github.com/hirokisan/design-pattern/code/structural/adapter/impl_1"
	"github.com/stretchr/testify/assert"
)

func TestAdapter(t *testing.T) {
	obj := adapter.NewImplement()

	assert.Equal(t, "implemented", obj.Method())
}
