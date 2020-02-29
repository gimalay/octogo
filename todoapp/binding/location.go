package binding

import (
	"github.com/golang/protobuf/proto"
)

func (loc *Location) mapToViewModel() viewModel {
	switch loc.Type {
	case LocationType_Home:
		return &ViewModel_Home{}
	case LocationType_Task:
		pl := &Location_Task{}
		err := proto.Unmarshal(loc.Payload, pl)
		if err != nil {
			panic(err)
		}
		return &ViewModel_Task{ID: pl.TaskID}
	case LocationType_Project:
		pl := &Location_Project{}
		err := proto.Unmarshal(loc.Payload, pl)
		if err != nil {
			panic(err)
		}
		return &ViewModel_Project{ID: pl.ProjectID}
	}
	return nil
}
