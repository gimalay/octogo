package model

import (
	"bytes"

	db "github.com/gimalay/binx"
	"github.com/golang/protobuf/proto"
)

type (
	ProjectsSlice []*Project
)

const (
	ProjectsBucket            = "ProjectsBucket"
	ProjectsMasterIndexBucket = "ProjectsMasterIndexBucket"
)

func (i *ProjectsSlice) BucketKey() []byte { return []byte(ProjectsBucket) }
func (i *ProjectsSlice) AppendBinary(data []byte) error {
	*i = append(*i, &Project{})
	return proto.Unmarshal(data, (*i)[len(*i)-1])
}

func (i *Project) Key() []byte                              { return i.ID }
func (i *Project) BucketKey() []byte                        { return []byte(ProjectsBucket) }
func (i *Project) MarshalBinary() ([]byte, error)           { return proto.Marshal(i) }
func (i *Project) UnmarshalBinary(jdata []byte) (err error) { return proto.Unmarshal(jdata, i) }

func (i *Project) Indexes() []db.Index {
	return []db.Index{}
}

func (i *Project) MasterIndexBucketKey() []byte { return []byte(ProjectsMasterIndexBucket) }

type ProjectEvent interface {
	proto.Message
	aggregate(i *Project, ev Event) (*Project, error)
}

func mapProjectEvent(t EventType) ProjectEvent {
	switch t {
	case EventType_ProjectCreated:
		return &Event_ProjectCreated{}
	case EventType_TaskRemoved:
		return &Event_TaskRemoved{}
	case EventType_TaskAdded:
		return &Event_TaskAdded{}
	case EventType_ProjectRenamed:
		return &Event_ProjectRenamed{}
	default:
		return nil
	}
}

func (i *Project) initEventType(et EventType) bool {
	return et == EventType_ProjectCreated
}

func (i *Project) aggregateEvent(ev Event) (aggregate, error) {
	ie := mapProjectEvent(ev.Type)
	if ie == nil {
		return i, errEventUnknownType
	}

	err := proto.Unmarshal(ev.Payload, ie)
	if err != nil {
		return i, err
	}

	if i != nil && i.initEventType(ev.Type) {
		return i, errProjectAlreadyCreated
	}
	if i == nil && !i.initEventType(ev.Type) {
		return i, errProjectDoesNotExist
	}

	return ie.aggregate(i, ev)
}

func (pl *Event_ProjectCreated) aggregate(i *Project, ev Event) (*Project, error) {
	i = &Project{}
	i.ID = ev.AggregateID
	i.Created = ev.Timestamp
	i.Updated = ev.Timestamp

	i.Name = pl.Name
	i.Tasks = pl.Tasks

	return i, nil
}

func (pl *Event_TaskRemoved) aggregate(i *Project, ev Event) (*Project, error) {
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

func (pl *Event_TaskAdded) aggregate(i *Project, ev Event) (*Project, error) {
	i.Tasks = append(i.Tasks, pl.TaskID)
	i.Updated = ev.Timestamp

	return i, nil
}

func (pl *Event_ProjectRenamed) aggregate(i *Project, ev Event) (*Project, error) {
	i.Name = pl.Name
	i.Updated = ev.Timestamp

	return i, nil
}
