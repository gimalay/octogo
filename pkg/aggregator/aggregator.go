package aggregator

import (
	"errors"
	"time"
)

type (
	EventWriter interface {
		WriteEvent(Event) error
	}
	EventReader interface {
		EventsForAggregate(aggregateID []byte) ([]Event, error)
	}
	AggregateWriter interface {
		WriteAggregate(Aggregate) error
	}
	Aggregate interface {
		AggregateEvent(Event) (Aggregate, error)
	}
	Command interface {
		Execute() ([]Event, error)
	}
	AggregateMapper interface {
		MapAggregate(Event) Aggregate
	}
)
type (
	Event struct {
		Type        int
		Payload     []byte
		AggregateID []byte
		Timestamp   time.Time
	}
	Aggregator struct {
		EventWriter
		EventReader
		AggregateWriter
		AggregateMapper
	}
)

func (e Event) validate() error {
	if e.Type == 0 {
		return errors.New("no event type")
	}

	if len(e.AggregateID) == 0 {
		return errors.New("no event aggregate ID")
	}
	if e.Timestamp.IsZero() {
		return errors.New("no event timestamp")
	}
	return nil
}

func (oc *Aggregator) Execute(cmd Command) error {
	events, err := cmd.Execute()

	if err != nil {
		return err
	}

	return oc.appendEvents(events)
}

func (oc *Aggregator) appendEvents(events []Event) error {
	for _, e := range events {
		err := e.validate()
		if err != nil {
			return err
		}
		err = oc.WriteEvent(e)
		if err != nil {
			return err
		}
		err = oc.updateAggregate(e.AggregateID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (oc *Aggregator) appendEvent(ag Aggregate, ev Event) (Aggregate, error) {
	if ag == nil {
		ag = oc.MapAggregate(ev)
	}
	return ag.AggregateEvent(ev)
}

func (oc *Aggregator) updateAggregate(id []byte) error {
	es, err := oc.EventReader.EventsForAggregate(id)
	if err != nil {
		return err
	}

	ag, err := oc.appendEvent(nil, es[0])
	if err != nil {
		return err
	}

	for _, e := range es[1:] {
		ag, err = oc.appendEvent(ag, e)
		if err != nil {
			return err
		}
	}

	return oc.AggregateWriter.WriteAggregate(ag)

}
