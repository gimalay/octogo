package aggregate

import (
	"github.com/gimalay/octogo/pkg/octogo"
	"github.com/golang/protobuf/proto"
)

func (ag *Task) AggregateEvent(e octogo.Event) (octogo.Aggregate, error) {
	h := ag.handler(EventType(e.Type))
	if h == nil {
		return ag, errEventUnknownType
	}

	err := proto.Unmarshal(e.Payload, h)
	if err != nil {
		return ag, err
	}

	ag, err = h.handle(ag, e)

	return ag, err
}

func (ag *Task) CanAggregate(ev octogo.Event) bool {
	return ag.handler(EventType(ev.Type)) != nil
}

type taskEventHandler interface {
	proto.Message
	handle(*Task, octogo.Event) (*Task, error)
}

func (i *Task) handler(t EventType) taskEventHandler {
	switch t {
	case EventType_TaskCreated:
		return &Task_Created{}
	case EventType_TaskRenamed:
		return &Task_Renamed{}
	default:
		return nil
	}
}

func (pl *Task_Renamed) handle(a *Task, ev octogo.Event) (*Task, error) {
	if a == nil {
		return a, errTaskDoesNotExist
	}

	a.Updated = ev.Timestamp

	a.Name = pl.Name

	return a, nil
}

func (pl *Task_Created) handle(a *Task, ev octogo.Event) (*Task, error) {
	if a != nil {
		return a, errTaskAlreadyCreated
	}

	a = &Task{}

	a.ID = ev.AggregateID
	a.Created = ev.Timestamp
	a.Updated = ev.Timestamp
	a.Name = pl.Name

	return a, nil
}
