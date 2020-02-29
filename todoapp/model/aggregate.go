package model

import (
	"sort"

	db "github.com/gimalay/binx"
	"github.com/golang/protobuf/ptypes"
)

type aggregate interface {
	db.Indexable
	aggregateEvent(Event) (aggregate, error)
}

func mapAggregateType(at AggregateType) aggregate {
	switch at {
	case AggregateType_Task:
		return (*Task)(nil)
	case AggregateType_Project:
		return (*Project)(nil)
	default:
		return nil
	}
}

func Aggregate(writer db.Writer, aggregateID []byte) error {
	events, err := eventsForAggregate(writer, aggregateID)
	if err != nil {
		return err
	}

	a := mapAggregateType(events[0].AggregateType)

	if a == nil {
		if events[0].SourceVersion > Version {
			//Aggregate from the future, skip aggregate
			return nil
		}
		return errUnknownAggregateType
	}

	for _, e := range events {
		a, err = a.aggregateEvent(e)

		if err == errEventUnknownType && e.SourceVersion > Version {
			//Event from the future, skip
			continue
		}

		if err != nil {
			return err
		}
	}

	if a != nil {
		return writer.Put(a)
	}

	return nil
}

func eventsForAggregate(reader db.Reader, aggregateID []byte) (EventSlice, error) {
	var events = EventSlice{}
	err := reader.List(&events, db.Where(EventAggregateIndex(aggregateID)))

	sort.Slice(events, func(i, j int) bool {
		it, err := ptypes.Timestamp(events[i].Timestamp)
		jt, err := ptypes.Timestamp(events[j].Timestamp)

		return err == nil && it.Before(jt) && jt.After(it) || events[i].ID < events[j].ID
	})

	if err != nil {
		return EventSlice{}, err
	}

	return events, err
}
