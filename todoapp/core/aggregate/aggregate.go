package aggregate

import (
	"github.com/gimalay/octogo/pkg/octogo"
)

type Mapper struct{}

func (a *Mapper) MapAggregate(ev octogo.Event) octogo.Aggregate {
	if ((*Task)(nil)).CanAggregate(ev) {
		return (*Task)(nil)
	}
	if ((*Project)(nil)).CanAggregate(ev) {
		return (*Project)(nil)
	}
	return nil
}
