package mediator_test

import (
	"testing"

	"github.com/hirokisan/design-pattern/code/behavioral/mediator"
	"github.com/stretchr/testify/assert"
)

func TestMediator(t *testing.T) {
	obj := mediator.NewMediator()

	obj.ColleagueA.SetEnabled(true)
	assert.True(t, obj.ColleagueB.Enabled())

	obj.ColleagueA.SetEnabled(false)
	assert.False(t, obj.ColleagueB.Enabled())
}
