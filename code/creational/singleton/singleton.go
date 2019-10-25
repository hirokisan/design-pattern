package singleton

import (
	"time"
)

type singleton struct {
	Now time.Time
}

func (s *singleton) Run() time.Time {
	return s.Now
}

var instance = newInstance()

func newInstance() *singleton {
	return &singleton{}
}

// NewInstance : このメソッドを介してインスタンスを生成することになるので、単一性が保たれるという話
func NewInstance() *singleton {
	return instance
}
