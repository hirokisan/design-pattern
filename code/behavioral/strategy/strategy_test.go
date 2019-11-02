package strategy_test

import (
	"testing"

	"github.com/hirokisan/design-pattern/code/behavioral/strategy"
	"github.com/stretchr/testify/assert"
)

func TestStrategy(t *testing.T) {
	objA := strategy.NewContext(strategy.NewStrategyA())
	objB := strategy.NewContext(strategy.NewStrategyB())

	assert.Equal(t, true, objA.Compare())
	assert.Equal(t, false, objB.Compare())
}
