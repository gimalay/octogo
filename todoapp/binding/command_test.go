package binding

import (
	"testing"

	db "github.com/gimalay/binx"
	"github.com/gimalay/octogo/todoapp/model"
	"github.com/stretchr/testify/assert"
)

func TestNewTask_Execute(t *testing.T) {
	tests := []struct {
		name               string
		existingTasks      []model.Task
		existingIdentities []model.Project
		argument           Command_NewTask
		expected           []newEvent
		err                error
	}{
		{
			name:               "new Task and Project created",
			existingIdentities: []model.Project{},
			existingTasks:      []model.Task{},
			argument: Command_NewTask{
				TaskID: []byte{1},
				Name:   "Task 1",
				Emoji:  ":)",
			},
			expected: []newEvent{
				newEvent{
					Type: model.EventType_TaskCreated,
					Payload: &model.Event_TaskCreated{
						Name:  "Task 1",
						Emoji: ":)",
					},
					AggregateID:   []byte{1},
					AggregateType: model.AggregateType_Task,
				},
			},
			err: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := commandContext{}
			d := &dbMock{}

			d.list = func(q db.QueryableSlice, cnd ...db.Condition) error { return nil }

			events, err := tt.argument.Execute(ctx, d)

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.expected, events)
		})
	}
}
