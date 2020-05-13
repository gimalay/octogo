package aggregate

import (
	"github.com/gimalay/octogo/pkg/aggregator"
)

type Mapper struct{}

func (a *Mapper) MapAggregate(ev aggregator.Event) aggregator.Aggregate {
	if ((*Task)(nil)).CanAggregate(ev) {
		return (*Task)(nil)
	}
	if ((*Project)(nil)).CanAggregate(ev) {
		return (*Project)(nil)
	}
	return nil
}
