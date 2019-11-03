package visitor_test

import (
	"testing"

	"github.com/hirokisan/design-pattern/code/behavioral/visitor"
	"github.com/stretchr/testify/assert"
)

func TestVisitor(t *testing.T) {
	v1 := visitor.NewVisitorA()
	v2 := visitor.NewVisitorB()

	elem := visitor.NewElement()

	assert.Equal(t, visitor.ErrFailed, elem.Accept(v1))
	assert.NoError(t, elem.Accept(v2))
}
