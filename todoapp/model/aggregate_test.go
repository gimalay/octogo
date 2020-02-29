package model

import (
	"testing"

	db "github.com/gimalay/binx"
	"github.com/stretchr/testify/assert"
)

func Test_Aggregate(t *testing.T) {
	tests := []struct {
		name     string
		existing []*Event
		argument []byte
		expected *Task
		err      error
	}{
		{
			name: "order and combine existing events by timestamp",
			existing: []*Event{
				&Event{
					ID:                  1,
					PreviousMessageHash: make([]byte, 32),
					AggregateID:         id(1),
					SourceID:            id(1),
					Payload:             bt(&Event_TaskRenamed{Name: "name2"}),
					Type:                EventType_TaskRenamed,
					AggregateType:       AggregateType_Task,
					Timestamp:           tm(2),
				},
				&Event{
					ID:                  2,
					PreviousMessageHash: make([]byte, 32),
					AggregateID:         id(1),
					SourceID:            id(1),
					Payload:             bt(&Event_TaskCreated{Name: "name1"}),
					Type:                EventType_TaskCreated,
					AggregateType:       AggregateType_Task,
					Timestamp:           tm(1),
				}},
			argument: id(1),
			expected: &Task{
				ID:      id(1),
				Name:    "name2",
				Created: tm(1),
				Updated: tm(2),
			},
		},
		{
			name: "order and combine existing with same date by ID",
			existing: []*Event{
				&Event{
					ID:                  2,
					PreviousMessageHash: make([]byte, 32),
					AggregateID:         id(1),
					SourceID:            id(1),
					Payload:             bt(&Event_TaskRenamed{Name: "name2"}),
					Type:                EventType_TaskRenamed,
					AggregateType:       AggregateType_Task,
					Timestamp:           tm(1),
				},
				&Event{
					ID:                  1,
					PreviousMessageHash: make([]byte, 32),
					AggregateID:         id(1),
					SourceID:            id(1),
					Payload:             bt(&Event_TaskCreated{Name: "name1"}),
					Type:                EventType_TaskCreated,
					AggregateType:       AggregateType_Task,
					Timestamp:           tm(1),
				}},
			argument: id(1),
			expected: &Task{
				ID:      id(1),
				Name:    "name2",
				Created: tm(1),
				Updated: tm(1),
			},
		},
		{
			name: "do not update entry if it cannot be aggregated",
			existing: []*Event{&Event{
				ID:                  1,
				PreviousMessageHash: make([]byte, 32),
				AggregateID:         id(1),
				SourceID:            id(1),
				Payload:             bt(&Event_TaskRenamed{Name: "name"}),
				Type:                EventType_TaskRenamed,
				AggregateType:       AggregateType_Task,
				Timestamp:           tm(1),
			}},
			argument: id(1),
			expected: nil,
			err:      errActivtiyDoesNotExist,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dm := &dbMock{}

			putTotal := 0
			dm.list = func(q db.QueryableSlice, cnd ...db.Condition) error {
				if tt.existing != nil {
					for _, v := range tt.existing {
						_ = q.AppendBinary(bt(v))
					}
				}
				return nil
			}
			dm.put = func(s db.Indexable) error {
				switch m := s.(type) {
				case *Event:
					assert.Equal(t, tt.argument, m.ID)
				case *Task:
					assertEqualTask(t, tt.expected, m)
				}
				putTotal++
				return nil
			}

			err := Aggregate(dm, tt.argument)

			assert.Equal(t, tt.err, err)
			if tt.expected != nil {
				assert.Equal(t, 1, putTotal)
			} else {
				assert.Equal(t, 0, putTotal)
			}
		})
	}
}
