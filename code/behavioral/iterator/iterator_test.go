package iterator_test

import (
	"testing"

	"github.com/hirokisan/design-pattern/code/behavioral/iterator"
	"github.com/stretchr/testify/assert"
)

func TestIterator(t *testing.T) {
	aggrList := iterator.NewAggr()
	obj1 := iterator.NewInst()
	obj2 := iterator.NewInst()

	aggrList.Add(obj1)
	aggrList.Add(obj2)
	iter := aggrList.Iterator()

	for iter.HasNext() {
		assert.True(t, iter.Next().Method())
	}
}
