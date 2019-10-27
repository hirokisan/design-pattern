package chain_of_responsibility_test

import (
	"testing"

	cor "github.com/hirokisan/design-pattern/code/behavioral/chain_of_responsibility"
	"github.com/stretchr/testify/assert"
)

func TestCOR(t *testing.T) {
	t.Run("staff", func(t *testing.T) {
		obj := cor.NewStaff()

		{
			staffIssue := cor.NewIssue(
				cor.IssueLevelStaff,
				"problem",
			)
			assert.Equal(t, "solved", obj.Handle(*staffIssue))
		}

		{
			managerIssue := cor.NewIssue(
				cor.IssueLevelManager,
				"problem",
			)
			assert.Equal(t, "solved", obj.Handle(*managerIssue))
		}

		{
			unknownIssue := cor.NewIssue(
				cor.IssueLevelUnknown,
				"problem",
			)
			assert.Equal(t, "not solved", obj.Handle(*unknownIssue))
		}
	})

	t.Run("manager", func(t *testing.T) {
		obj := cor.NewManager()

		{
			staffIssue := cor.NewIssue(
				cor.IssueLevelStaff,
				"problem",
			)
			assert.Equal(t, "not solved", obj.Handle(*staffIssue))
		}

		{
			managerIssue := cor.NewIssue(
				cor.IssueLevelManager,
				"problem",
			)
			assert.Equal(t, "solved", obj.Handle(*managerIssue))
		}

		{
			unknownIssue := cor.NewIssue(
				cor.IssueLevelUnknown,
				"problem",
			)
			assert.Equal(t, "not solved", obj.Handle(*unknownIssue))
		}
	})
}
