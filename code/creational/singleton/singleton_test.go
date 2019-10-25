package singleton_test

import (
	"testing"

	"github.com/hirokisan/design-pattern/code/creational/singleton"
	"github.com/stretchr/testify/assert"
)

func TestSingleton(t *testing.T) {
	obj := singleton.NewInstance()
	obj2 := singleton.NewInstance()

	assert.Exactly(t, obj, obj2)
	assert.Equal(t, obj.Run(), obj2.Run())
}

// TODO : implement
//func TestSingletonWithMultiThread(t *testing) {}
