package state_test

import (
	"testing"

	"github.com/hirokisan/design-pattern/code/behavioral/state"
	"github.com/stretchr/testify/assert"
)

func TestState(t *testing.T) {
	obj := state.NewPerson()
	assert.Equal(t, "happy", obj.ExpressFeeling())

	obj.SetFeeling(state.FeelingAngry)
	assert.Equal(t, "angry", obj.ExpressFeeling())
}
