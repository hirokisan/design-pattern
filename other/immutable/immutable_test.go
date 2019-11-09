package immutable_test

import (
	"testing"

	"github.com/hirokisan/design-pattern/other/immutable"
	"github.com/stretchr/testify/assert"
)

func TestImmutable(t *testing.T) {
	obj := immutable.NewImmutable("immutable")

	assert.Equal(t, "immutable", obj.Name(), "read only")
}
