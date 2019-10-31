package iterator

type Aggregation interface {
	Iterator() Iterator
}

// Aggr : 集約オブジェクト
type Aggr struct {
	list []Instance
}

func NewAggr() *Aggr {
	return &Aggr{}
}

func (a *Aggr) Iterator() Iterator {
	// newAggrIterは他のiteratorを返しても良い
	return newAggrIter(a)
}

// Add : interfaceには定義されていない
func (a *Aggr) Add(instance Instance) {
	a.list = append(a.list, instance)
}
