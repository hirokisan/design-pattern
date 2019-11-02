package strategy

type strategy interface {
	compare() bool
}

type strategyA struct {
}

func NewStrategyA() strategy {
	return &strategyA{}
}

func (s *strategyA) compare() bool {
	return true
}

type strategyB struct {
}

func NewStrategyB() strategy {
	return &strategyB{}
}

func (s *strategyB) compare() bool {
	return false
}

type context struct {
	strategy strategy
}

func NewContext(strategy strategy) *context {
	return &context{
		strategy: strategy,
	}
}

func (c *context) Compare() bool {
	return c.strategy.compare()
}
