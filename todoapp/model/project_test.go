package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProject_Aggregate(t *testing.T) {
	tests := []struct {
		name     string
		existing aggregate
		argument Event
		expected *Project
		err      error
	}{
		{
			name:     "Project created",
			existing: (*Project)(nil),
			argument: Event{
				AggregateID: id(1),
				Payload: bt(&Event_ProjectCreated{
					Name:  "name1",
					Tasks: [][]byte{[]byte{1}},
				}),
				Type:          EventType_ProjectCreated,
				AggregateType: AggregateType_Project,
				Timestamp:     tm(1),
			},
			expected: &Project{
				ID:      id(1),
				Name:    "name1",
				Tasks:   [][]byte{[]byte{1}},
				Created: tm(1),
				Updated: tm(1),
			},
			err: nil,
		},
		{
			name: "remove one of three Tasks",
			existing: &Project{
				Tasks: [][]byte{
					[]byte{1},
					[]byte{2},
					[]byte{3},
				},
				Created: tm(1),
			},
			argument: Event{
				AggregateID: id(1),
				Payload: bt(&Event_TaskRemoved{
					TaskID: []byte{2},
				}),
				Type:          EventType_TaskRemoved,
				AggregateType: AggregateType_Project,
				Timestamp:     tm(1),
			},
			expected: &Project{
				Tasks: [][]byte{
					[]byte{1},
					[]byte{3},
				},
				Created: tm(1),
				Updated: tm(1),
			},
			err: nil,
		},
		{
			name: "remove Task",
			existing: &Project{
				Tasks:   [][]byte{[]byte{1}},
				Created: tm(1),
			},
			argument: Event{
				AggregateID: id(1),
				Payload: bt(&Event_TaskRemoved{
					TaskID: []byte{1},
				}),
				Type:          EventType_TaskRemoved,
				AggregateType: AggregateType_Project,
				Timestamp:     tm(1),
			},
			expected: &Project{
				Tasks:   [][]byte{},
				Created: tm(1),
				Updated: tm(1),
			},
			err: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entry, err := tt.existing.aggregateEvent(tt.argument)

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.expected, entry)
		})
	}
}
