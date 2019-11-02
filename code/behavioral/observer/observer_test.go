package observer_test

import (
	"testing"

	"github.com/hirokisan/design-pattern/code/behavioral/observer"
	"github.com/stretchr/testify/assert"
)

func TestObserver(t *testing.T) {
	obs := observer.NewObserver()
	sub := observer.NewSubject("tom")
	sub.AddObserver(obs)

	assert.NoError(t, sub.ChangeName("subject"))
	assert.Equal(t, observer.ErrNofityObserver, sub.ChangeName("observer"))
}
