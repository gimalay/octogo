package binding

import (
	fmt "fmt"

	db "github.com/gimalay/binx"
	"github.com/gimalay/octogo/todoapp/model"
)

func mapCommandPayload(t CommandType) command {
	switch t {
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
	case CommandType_DeleteProject:
		return &Command_DeleteProject{}
	}
	panic(fmt.Sprintf("unknown command type: %v", t))
}

func (c Command_NewTask) Execute(ctx commandContext, r db.Reader) ([]newEvent, error) {
	events := []newEvent{
		newEvent{
			Type: model.EventType_TaskCreated,
			Payload: &model.Event_TaskCreated{
				Name:  c.Name,
				Emoji: c.Emoji,
			},
			AggregateID:   c.TaskID,
			AggregateType: model.AggregateType_Task,
		},
	}

	return events, nil
}

func (c Command_NewProject) Execute(ctx commandContext, r db.Reader) ([]newEvent, error) {
	events := []newEvent{
		newEvent{
			Type: model.EventType_ProjectCreated,
			Payload: &model.Event_ProjectCreated{
				Name:  c.Name,
				Tasks: c.Tasks,
			},
			AggregateID:   c.ProjectID,
			AggregateType: model.AggregateType_Project,
		},
	}

	return events, nil
}

func (c Command_RemoveTask) Execute(ctx commandContext, r db.Reader) ([]newEvent, error) {
	return []newEvent{
		newEvent{
			Type: model.EventType_TaskRemoved,
			Payload: &model.Event_TaskRemoved{
				TaskID: c.TaskID,
			},
			AggregateID:   c.ProjectID,
			AggregateType: model.AggregateType_Project,
		},
	}, nil
}

func (c Command_AddTask) Execute(ctx commandContext, r db.Reader) ([]newEvent, error) {
	return []newEvent{
		newEvent{
			Type: model.EventType_TaskAdded,
			Payload: &model.Event_TaskAdded{
				TaskID: c.TaskID,
			},
			AggregateID:   c.ProjectID,
			AggregateType: model.AggregateType_Project,
		},
	}, nil
}

func (c Command_RenameTask) Execute(ctx commandContext, r db.Reader) ([]newEvent, error) {
	return []newEvent{
		newEvent{
			Type: model.EventType_TaskRenamed,
			Payload: &model.Event_TaskRenamed{
				Name:  c.Name,
				Emoji: c.Emoji,
			},
			AggregateID:   c.TaskID,
			AggregateType: model.AggregateType_Task,
		},
	}, nil
}

func (c Command_DeleteProject) Execute(ctx commandContext, r db.Reader) ([]newEvent, error) {
	return []newEvent{
		newEvent{
			Type:          model.EventType_ProjectDeleted,
			Payload:       &model.Event_ProjectDeleted{},
			AggregateID:   c.ProjectID,
			AggregateType: model.AggregateType_Project,
		},
	}, nil
}

func (c Command_RenameProject) Execute(ctx commandContext, r db.Reader) ([]newEvent, error) {
	return []newEvent{
		newEvent{
			Type: model.EventType_ProjectRenamed,
			Payload: &model.Event_ProjectRenamed{
				Name: c.Name,
			},
			AggregateID:   c.ProjectID,
			AggregateType: model.AggregateType_Project,
		},
	}, nil
}
