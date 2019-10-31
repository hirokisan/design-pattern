package iterator

type Iterator interface {
	HasNext() bool
	Next() Instance
}

type aggrIter struct {
	index int
	aggr  *Aggr
}

func newAggrIter(aggr *Aggr) *aggrIter {
	return &aggrIter{
		aggr: aggr,
	}
}

func (i *aggrIter) HasNext() bool {
	if len(i.aggr.list) > i.index {
		return true
	}

	return false
}

func (i *aggrIter) Next() Instance {
	instance := i.aggr.list[i.index]
	i.index++
	return instance
}
