package model

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestTask_Aggregate(t *testing.T) {
	tests := []struct {
		name     string
		existing aggregate
		argument Event
		expected *Task
		err      error
	}{
		{
			name:     "Task created",
			existing: (*Task)(nil),
			argument: Event{
				AggregateID: id(1),
				Payload: bt(&Event_TaskCreated{
					Name:  "name1",
					Emoji: "1",
				}),
				Type:          EventType_TaskCreated,
				AggregateType: AggregateType_Task,
				Timestamp:     tm(1),
			},
			expected: &Task{
				ID:      id(1),
				Name:    "name1",
				Emoji:   "1",
				Created: tm(1),
				Updated: tm(1),
			},
			err: nil,
		},
		{
			name: "Task already created error",
			existing: &Task{
				ID: id(1),
			},
			argument: Event{
				AggregateID:   id(1),
				Payload:       bt(&Event_TaskCreated{}),
				Type:          EventType_TaskCreated,
				AggregateType: AggregateType_Task,
				Timestamp:     tm(1),
			},
			expected: &Task{
				ID: id(1),
			},
			err: errTaskAlreadyCreated,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entry, err := tt.existing.aggregateEvent(tt.argument)
			assert.Equal(t, err, tt.err)

			if tt.expected != nil && entry != nil {
				assert.Empty(t, cmp.Diff(tt.expected, entry))
			} else {
				assert.Empty(t, tt.expected, entry)
			}

		})
	}
}
