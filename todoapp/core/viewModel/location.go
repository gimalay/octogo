package viewModel

import (
	"errors"
	"github.com/gimalay/octogo/todoapp/core/aggregate"
	"time"

	"github.com/golang/protobuf/proto"
)

type Reader interface {
	Tasks() ([]*aggregate.Task, error)
	Task(id []byte) (*aggregate.Task, error)
	Project(id []byte) (*aggregate.Project, error)
	Projects() ([]*aggregate.Project, error)
}

type viewModel interface {
	Read(Reader) error
	proto.Message
}

type Message struct {
	Type      int
	Payload   []byte
	Timestamp time.Time
}

type Navigator struct {
	Reader
}

func (loc Message) mapLocation() viewModel {
	switch LocationType(loc.Type) {
	case LocationType_Home:
		return &ViewModel_Home{}
	case LocationType_Project:
		pl := &Location_Project{}
		err := proto.Unmarshal(loc.Payload, pl)
		if err != nil {
			panic(err)
		}
		return &ViewModel_Project{ID: pl.ProjectID}
	case LocationType_Task:
		pl := &Location_Task{}
		err := proto.Unmarshal(loc.Payload, pl)
		if err != nil {
			panic(err)
		}
		return &ViewModel_Task{ID: pl.TaskID}
	default:
		return nil
	}
}

func (n *Navigator) NavigateTo(message Message) ([]byte, error) {

	vm := message.mapLocation()

	if vm == nil {
		return []byte{}, errors.New("cannot define View Model type")
	}

	err := vm.Read(n.Reader)

	if err != nil {
		return []byte{}, err
	}

	return proto.Marshal(vm)
}
