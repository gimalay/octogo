package aggregate

import (
	"github.com/gimalay/octogo/pkg/aggregator"
	"testing"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTask_Aggregate(t *testing.T) {
	tests := []struct {
		name     string
		existing aggregator.Aggregate
		argument aggregator.Event
		expected *Task
		err      error
	}{
		{
			name:     "task created",
			existing: (*Task)(nil),
			argument: aggregator.Event{
				AggregateID: id(1),
				Payload: bt(&Task_Created{
					Name: "name1",
				}),
				Type:      int(EventType_TaskCreated),
				Timestamp: tm(1),
			},
			expected: &Task{
				ID:      id(1),
				Name:    "name1",
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
			argument: aggregator.Event{
				AggregateID: id(1),
				Payload:     bt(&Task_Created{}),
				Type:        int(EventType_TaskCreated),
				Timestamp:   tm(1),
			},
			expected: &Task{
				ID: id(1),
			},
			err: errTaskAlreadyCreated,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ag, err := tt.existing.AggregateEvent(tt.argument)
			assert.Equal(t, err, tt.err)
			assert.Equal(t, tt.expected, ag)
		})
	}
}
func id(b byte) []byte {
	v, _ := uuid.NewMD5(uuid.NameSpaceDNS, []byte{b}).MarshalBinary()
	return v

}
func tm(s time.Duration) time.Time {
	return time.Time{}.Add(s)
}

func ttm(s int64) time.Time {
	return time.Time{}.Add(time.Second * time.Duration(s))
}

func bt(message proto.Message) []byte {
	b, _ := proto.Marshal(message)
	return b
}
