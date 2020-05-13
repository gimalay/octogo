package event

import (
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestEventType_validate(t *testing.T) {
	tests := []struct {
		name     string
		existing *Event
		err      error
	}{
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
				Type:                1,
				Timestamp:           time.Time{}.Add(1),
			},
			nil,
		},
		{
			"no timestamp",
			&Event{
				ID:                  1,
				PreviousMessageHash: make([]byte, 32),
				SourceID:            id(2),
				AggregateID:         id(3),
				Type:                1,
				Timestamp:           time.Time{},
			},
			errNoTimestamp,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.existing.Validate()

			assert.Equal(t, tt.err, err)
		})
	}
}
func id(b byte) []byte {
	v, _ := uuid.NewMD5(uuid.NameSpaceDNS, []byte{b}).MarshalBinary()
	return v

}
func tm(s int64) *timestamp.Timestamp {
	r, _ := ptypes.TimestampProto(ttm(s))
	return r
}

func ttm(s int64) time.Time {
	return time.Time{}.Add(time.Second * time.Duration(s))
}

func bt(message proto.Message) []byte {
	b, _ := proto.Marshal(message)
	return b
}
