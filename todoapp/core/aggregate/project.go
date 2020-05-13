package aggregate

import (
	"bytes"
	"github.com/gimalay/octogo/pkg/aggregator"
	"github.com/gogo/protobuf/proto"
)

func (ag *Project) AggregateEvent(e aggregator.Event) (aggregator.Aggregate, error) {
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

func (ag *Project) CanAggregate(ev aggregator.Event) bool {
	return ag.handler(EventType(ev.Type)) != nil
}

type ProjectEventHandler interface {
	proto.Message
	handle(*Project, aggregator.Event) (*Project, error)
}

func (i *Project) handler(t EventType) ProjectEventHandler {
	switch t {
	case EventType_ProjectCreated:
		return &Project_Created{}
	case EventType_TaskRemoved:
		return &Project_TaskRemoved{}
	case EventType_TaskAdded:
		return &Project_TaskAdded{}
	case EventType_ProjectRenamed:
		return &Project_Renamed{}
	default:
		return nil
	}
}

func (pl *Project_Created) handle(i *Project, ev aggregator.Event) (*Project, error) {
	if i != nil {
		return nil, errProjectAlreadyCreated
	}

	i = &Project{}

	i.ID = ev.AggregateID
	i.Created = ev.Timestamp
	i.Updated = ev.Timestamp

	i.Name = pl.Name
	i.Tasks = pl.Tasks

	return i, nil
}

func (pl *Project_TaskRemoved) handle(i *Project, ev aggregator.Event) (*Project, error) {
	if i == nil {
		return nil, errProjectDoesNotExist
	}
	Tasks := [][]byte{}

	for _, a := range i.Tasks {
		if !bytes.Equal(a, pl.TaskID) {
			Tasks = append(Tasks, a)
		}
	}

	if len(i.Tasks) != len(Tasks) {
		i.Tasks = Tasks
		i.Updated = ev.Timestamp
	}
	return i, nil
}

func (pl *Project_TaskAdded) handle(i *Project, ev aggregator.Event) (*Project, error) {
	if i == nil {
		return nil, errProjectDoesNotExist
	}
	i.Tasks = append(i.Tasks, pl.TaskID)
	i.Updated = ev.Timestamp

	return i, nil
}

func (pl *Project_Renamed) handle(i *Project, ev aggregator.Event) (*Project, error) {
	if i == nil {
		return nil, errProjectDoesNotExist
	}
	i.Name = pl.Name
	i.Updated = ev.Timestamp

	return i, nil
}
