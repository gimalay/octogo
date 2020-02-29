package model

import (
	"testing"

	db "github.com/gimalay/binx"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/stretchr/testify/assert"
)

func TestEventType_validate(t *testing.T) {
	tests := []struct {
		name     string
		existing *Event
		err      error
	}{
		{
			"event 0 is invalid",
			&Event{
				ID:                  0,
				PreviousMessageHash: make([]byte, 32),
				SourceID:            id(1),
				AggregateID:         id(3),
			},
			errCannotValidateEvent0,
		},
		{
			"event 1 has zeros previous hash",
			&Event{
				ID:                  1,
				PreviousMessageHash: eventHash(&Event{ID: 2}),
				SourceID:            id(1),
				AggregateID:         id(3),
				Type:                EventType_TaskCreated,
				Timestamp:           &timestamp.Timestamp{},
				AggregateType:       AggregateType_Task,
			},
			errSpecialPrevMessageHash,
		},
		{
			"previous message hash has unexpected len",
			&Event{
				ID:                  1,
				PreviousMessageHash: make([]byte, 1),
				SourceID:            id(1),
				AggregateID:         id(3),
			},
			errHashLen,
		},
		{
			"empty source id",
			&Event{
				ID:                  1,
				PreviousMessageHash: make([]byte, 32),
				SourceID:            []byte{},
				AggregateID:         id(3),
			},
			errSourceID,
		},
		{
			"non guid source id",
			&Event{
				ID:                  1,
				PreviousMessageHash: make([]byte, 32),
				SourceID:            []byte{1},
				AggregateID:         id(3),
			},
			errSourceID,
		},
		{
			"empty previous event id",
			&Event{
				ID:                  1,
				PreviousMessageHash: make([]byte, 32),
				SourceID:            id(2),
				AggregateID:         id(3),
				Type:                EventType_TaskCreated,
				Timestamp:           &timestamp.Timestamp{},
				AggregateType:       AggregateType_Task,
			},
			nil,
		},
		{
			"unknown event type",
			&Event{
				ID:                  1,
				PreviousMessageHash: make([]byte, 32),
				SourceID:            id(2),
				AggregateID:         id(3),
				Type:                EventType_Unknown,
			},
			errUnknownType,
		},
		{
			"no timestamp",
			&Event{
				ID:                  1,
				PreviousMessageHash: make([]byte, 32),
				SourceID:            id(2),
				AggregateID:         id(3),
				Type:                EventType_TaskCreated,
				AggregateType:       AggregateType_Task,
				Timestamp:           nil,
			},
			errNoTimestamp,
		},
		{
			"unknown aggregate type",
			&Event{
				ID:                  1,
				PreviousMessageHash: make([]byte, 32),
				SourceID:            id(2),
				AggregateID:         id(3),
				Type:                EventType_TaskCreated,
				Timestamp:           &timestamp.Timestamp{},
				AggregateType:       AggregateType_UnknownAggregate,
			},
			errUnknownAggregateType,
		},
		{
			"event type and aggregate type missmatch",
			&Event{
				ID:                  1,
				PreviousMessageHash: make([]byte, 32),
				SourceID:            id(2),
				AggregateID:         id(3),
				Type:                EventType_TaskCreated,
				Timestamp:           &timestamp.Timestamp{},
				AggregateType:       AggregateType_Project,
			},
			errIncompatibleEvent,
		},
		{
			"unknown event type ignored if source version is bigger then current",
			&Event{
				ID:                  1,
				PreviousMessageHash: make([]byte, 32),
				SourceID:            id(2),
				AggregateID:         id(3),
				Type:                EventType(1231), //1231 = uknown event type
				SourceVersion:       1000,
				Timestamp:           &timestamp.Timestamp{},
				AggregateType:       AggregateType_Project,
			},
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.existing.Validate()

			assert.Equal(t, tt.err, err)
		})
	}
}

func TestAppendEvent(t *testing.T) {
	tests := []struct {
		name     string
		existing *Event
		argument *Event
		expected *Task
		err      error
	}{
		{
			name:     "returns error if event is invalid",
			existing: nil,
			argument: &Event{
				ID:                  1,
				PreviousMessageHash: make([]byte, 32),
				AggregateID:         id(1),
				SourceID:            id(1),
				Timestamp:           tm(1),
			},
			err: errUnknownType,
		},
		{
			name:     "append event if it is valid and has a matching hash and ID is previous ID + 1",
			existing: &Event{ID: 1},
			argument: &Event{
				ID:                  2,
				PreviousMessageHash: eventHash(&Event{ID: 1}),
				AggregateID:         id(1),
				AggregateType:       AggregateType_Task,
				SourceID:            id(1),
				Type:                EventType_TaskCreated,
				Timestamp:           tm(1),
			},
			err: nil,
		},
		{
			name:     "return error is ID is not squential",
			existing: &Event{ID: 1},
			argument: &Event{
				ID:                  3,
				PreviousMessageHash: eventHash(&Event{ID: 1}),
				AggregateID:         id(1),
				AggregateType:       AggregateType_Task,
				SourceID:            id(1),
				Type:                EventType_TaskCreated,
				Timestamp:           tm(1),
			},
			err: errUnexpectedID,
		},

		{
			name:     "append event with no hash for previous event",
			existing: &Event{ID: 2},
			argument: &Event{
				ID:                  3,
				PreviousMessageHash: make([]byte, 32),
				AggregateID:         id(1),
				AggregateType:       AggregateType_Task,
				SourceID:            id(1),
				Type:                EventType_TaskCreated,
				Timestamp:           tm(1),
			},
			err: errPrevEventHash,
		},
		{
			name:     "return error if previous event hash is diffrent",
			existing: &Event{ID: 2},
			argument: &Event{
				ID:                  3,
				PreviousMessageHash: eventHash(&Event{ID: 4}),
				AggregateID:         id(1),
				AggregateType:       AggregateType_Task,
				SourceID:            id(1),
				Type:                EventType_TaskCreated,
				Timestamp:           tm(1),
			},
			err: errPrevEventHash,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dm := &dbMock{}
			dm.put = func(s db.Indexable) error {
				return nil
			}

			err := AppendEvent(dm, tt.argument, tt.existing)
			if tt.err != nil {
				assert.Equal(t, tt.err, err)
				assert.Equal(t, 0, dm.putCallsCount)
			} else {
				assert.Equal(t, 1, dm.putCallsCount)
			}
		})
	}
}

func eventHash(e *Event) []byte {
	hash, _ := e.Hash()
	return hash[:]
}

func assertEqualTask(t *testing.T, expected, actual *Task) {
	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.Emoji, actual.Emoji)
	assert.Equal(t, expected.Updated, actual.Updated)
	assert.Equal(t, expected.Created, actual.Created)
}
