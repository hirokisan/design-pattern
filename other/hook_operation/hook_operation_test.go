package hook_operation_test

import (
	"errors"
	"testing"

	"github.com/hirokisan/design-pattern/other/hook_operation"
	"github.com/stretchr/testify/assert"
)

func TestHookOperation(t *testing.T) {
	t.Run("none, no error", func(t *testing.T) {
		obj := hook_operation.NewObject()

		assert.NoError(t, obj.Run())
	})
	t.Run("pre, no error", func(t *testing.T) {
		preHook := func() error {
			return nil
		}
		obj := hook_operation.NewObject().
			WithPreHooks(preHook)

		assert.NoError(t, obj.Run())
	})
	t.Run("pre and post, no error", func(t *testing.T) {
		preHook := func() error {
			return nil
		}
		postHook := func() error {
			return nil
		}
		obj := hook_operation.NewObject().
			WithPreHooks(preHook).
			WithPostHooks(postHook)

		assert.NoError(t, obj.Run())
	})
	t.Run("pre, error", func(t *testing.T) {
		preHook := func() error {
			return errors.New("error")
		}
		obj := hook_operation.NewObject().
			WithPreHooks(preHook)

		assert.Error(t, obj.Run())
	})
}
