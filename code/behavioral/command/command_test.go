package command_test

import (
	"testing"

	"github.com/hirokisan/design-pattern/code/behavioral/command"
	"github.com/stretchr/testify/assert"
)

func TestCommand(t *testing.T) {
	cmd := command.NewCommand()
	res := cmd.Execute()

	assert.True(t, res)
}
