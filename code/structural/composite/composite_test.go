package composite_test

import (
	"testing"

	"github.com/hirokisan/design-pattern/code/structural/composite"
)

func TestComposite(t *testing.T) {
	component := composite.NewPerson("taro")
	component.Greet()

	composite := composite.NewComposite()
	composite.Add(component)
	composite.Greet()
}
