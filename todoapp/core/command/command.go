package command

import (
	fmt "fmt"
	"github.com/gimalay/octogo/pkg/octogo"
	"github.com/gimalay/octogo/todoapp/core/aggregate"
	"time"

	"github.com/golang/protobuf/proto"
)

type Cmd interface {
	Execute(Reader) ([]octogo.Event, error)
	proto.Message
}

type Reader interface {
	Project([]byte) (*aggregate.Project, error)
	Task([]byte) (*aggregate.Task, error)
}

type commandAdapter struct {
	Reader
	Cmd
	timestamp time.Time
}

func (c *commandAdapter) Execute() ([]octogo.Event, error) {
	ev, err := c.Cmd.Execute(c.Reader)

	r := []octogo.Event{}

	for _, e := range ev {
		e.Timestamp = c.timestamp
		r = append(r, e)
	}

	return r, err
}

type Unmarshaler struct {
	Reader
}

type Message struct {
	Type      int
	Payload   []byte
	Timestamp time.Time
}

func (u *Unmarshaler) Unmarshal(data Message) (octogo.Command, error) {
	c := MapCommandPayload(data.Type)
	err := proto.Unmarshal(data.Payload, c)
	if err != nil {
		panic("cannot unmarshal")
	}
	return &commandAdapter{Cmd: c, Reader: u.Reader, timestamp: data.Timestamp}, nil
}

func MapCommandPayload(t int) Cmd {
	switch CommandType(t) {
	case CommandType_NewTask:
		return &Command_NewTask{}
	case CommandType_NewProject:
		return &Command_NewProject{}
	case CommandType_RemoveTask:
		return &Command_RemoveTask{}
	case CommandType_AddTask:
		return &Command_AddTask{}
	case CommandType_RenameTask:
		return &Command_RenameTask{}
	case CommandType_RenameProject:
		return &Command_RenameProject{}
	}
	panic(fmt.Sprintf("unknown command type: %v", t))
}

func (c Command_NewTask) Execute(r Reader) ([]octogo.Event, error) {
	events := []octogo.Event{
		{
			Type: int(aggregate.EventType_TaskCreated),
			Payload: bytes(&aggregate.Task_Created{
				Name: c.Name,
			}),
			AggregateID: c.TaskID,
		},
	}

	return events, nil
}

func (c Command_NewProject) Execute(r Reader) ([]octogo.Event, error) {
	events := []octogo.Event{
		{
			Type: int(aggregate.EventType_ProjectCreated),
			Payload: bytes(&aggregate.Project_Created{
				Name:  c.Name,
				Tasks: c.Tasks,
			}),
			AggregateID: c.ProjectID,
		},
	}

	return events, nil
}

func (c Command_RemoveTask) Execute(r Reader) ([]octogo.Event, error) {
	return []octogo.Event{
		{
			Type: int(aggregate.EventType_TaskRemoved),
			Payload: bytes(&aggregate.Project_TaskRemoved{
				TaskID: c.TaskID,
			}),
			AggregateID: c.ProjectID,
		},
	}, nil
}

func (c Command_AddTask) Execute(r Reader) ([]octogo.Event, error) {
	return []octogo.Event{
		{
			Type: int(aggregate.EventType_TaskAdded),
			Payload: bytes(&aggregate.Project_TaskAdded{
				TaskID: c.TaskID,
			}),
			AggregateID: c.ProjectID,
		},
	}, nil
}

func (c Command_RenameTask) Execute(r Reader) ([]octogo.Event, error) {
	return []octogo.Event{
		{
			Type: int(aggregate.EventType_TaskRenamed),
			Payload: bytes(&aggregate.Task_Renamed{
				Name: c.Name,
			}),
			AggregateID: c.TaskID,
		},
	}, nil
}
func (c Command_RenameProject) Execute(r Reader) ([]octogo.Event, error) {
	return []octogo.Event{
		{
			Type: int(aggregate.EventType_ProjectRenamed),
			Payload: bytes(&aggregate.Project_Renamed{
				Name: c.Name,
			}),
			AggregateID: c.ProjectID,
		},
	}, nil
}

func bytes(m proto.Message) []byte {
	b, err := proto.Marshal(m)
	if err != nil {
		panic(err)
	}
	return b
}
