package model

import (
	db "github.com/gimalay/binx"
	"github.com/golang/protobuf/proto"
)

type (
	TasksSlice []*Task
)

const (
	TasksBucket            = "TasksBucket"
	TasksMasterIndexBucket = "TasksMasterIndexBucket"
)

func (a *TasksSlice) BucketKey() []byte { return []byte(TasksBucket) }
func (a *TasksSlice) AppendBinary(data []byte) error {
	*a = append(*a, &Task{})
	return proto.Unmarshal(data, (*a)[len(*a)-1])
}

func (a *Task) Key() []byte                              { return a.ID }
func (a *Task) BucketKey() []byte                        { return []byte(TasksBucket) }
func (a *Task) MarshalBinary() ([]byte, error)           { return proto.Marshal(a) }
func (a *Task) UnmarshalBinary(jdata []byte) (err error) { return proto.Unmarshal(jdata, a) }

func (a *Task) Indexes() []db.Index {
	return []db.Index{}
}
func (a *Task) MasterIndexBucketKey() []byte { return []byte(TasksMasterIndexBucket) }

type TaskEvent interface {
	proto.Message
	aggregate(i *Task, ev Event) (*Task, error)
}

func mapTaskEvent(t EventType) TaskEvent {
	switch t {
	case EventType_TaskCreated:
		return &Event_TaskCreated{}
	case EventType_TaskRenamed:
		return &Event_TaskRenamed{}
	default:
		return nil
	}
}

func (i *Task) initEventType(et EventType) bool {
	return et == EventType_TaskCreated || et == EventType_TaskImported
}

func (i *Task) aggregateEvent(ev Event) (aggregate, error) {
	ie := mapTaskEvent(ev.Type)

	if ie == nil {
		return i, errEventUnknownType
	}

	err := proto.Unmarshal(ev.Payload, ie)
	if err != nil {
		return i, err
	}

	if i != nil && i.initEventType(ev.Type) {
		return i, errTaskAlreadyCreated
	}
	if i == nil && !i.initEventType(ev.Type) {
		return i, errActivtiyDoesNotExist
	}
	return ie.aggregate(i, ev)
}

func (pl *Event_TaskRenamed) aggregate(a *Task, ev Event) (*Task, error) {
	a.Updated = ev.Timestamp

	a.Name = pl.Name
	a.Emoji = pl.Emoji

	return a, nil
}

func (pl *Event_TaskCreated) aggregate(a *Task, ev Event) (*Task, error) {
	a = &Task{}
	a.ID = ev.AggregateID
	a.Created = ev.Timestamp
	a.Updated = ev.Timestamp
	a.Name = pl.Name
	a.Emoji = pl.Emoji

	return a, nil
}
