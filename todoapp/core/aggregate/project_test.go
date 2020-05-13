package aggregate

import (
	"github.com/gimalay/octogo/pkg/aggregator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProject_Aggregate(t *testing.T) {
	tests := []struct {
		name     string
		existing aggregator.Aggregate
		argument aggregator.Event
		expected *Project
		err      error
	}{
		{
			name:     "Project created",
			existing: (*Project)(nil),
			argument: aggregator.Event{
				AggregateID: id(1),
				Payload: bt(&Project_Created{
					Name:  "name1",
					Tasks: [][]byte{[]byte{1}},
				}),
				Type:      int(EventType_ProjectCreated),
				Timestamp: tm(1),
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
			argument: aggregator.Event{
				AggregateID: id(1),
				Payload: bt(&Project_TaskRemoved{
					TaskID: []byte{2},
				}),
				Type:      int(EventType_TaskRemoved),
				Timestamp: tm(1),
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
			name: "remove a Task",
			existing: &Project{
				Tasks:   [][]byte{[]byte{1}},
				Created: tm(1),
			},
			argument: aggregator.Event{
				AggregateID: id(1),
				Payload: bt(&Project_TaskRemoved{
					TaskID: []byte{1},
				}),
				Type:      int(EventType_TaskRemoved),
				Timestamp: tm(1),
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
			ag, err := tt.existing.AggregateEvent(tt.argument)
			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.expected, ag)
		})
	}
}
