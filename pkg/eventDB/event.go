package eventDB

import (
	"encoding/binary"

	db "github.com/gimalay/binx"
	"github.com/gimalay/octogo/pkg/event"

	"github.com/golang/protobuf/proto"
)

type eventIndexable event.Event

type (
	EventSlice          []*event.Event
	aggregateIndex      []byte
	EventIDReverseIndex uint64
)

const eventsBucket = "EventsBucket"

func (s *EventSlice) BucketKey() []byte { return []byte(eventsBucket) }
func (s *EventSlice) AppendBinary(data []byte) (bool, error) {
	*s = append(*s, &event.Event{})
	return true, proto.Unmarshal(data, (*s)[len(*s)-1])
}

func (e *eventIndexable) UniqueKey() []byte {
	return EventKey(e.ID)
}
func (e *eventIndexable) BucketKey() []byte { return []byte(eventsBucket) }
func (e *eventIndexable) MarshalBinary() ([]byte, error) {
	return proto.Marshal((*event.Event)(e))
}
func (e *eventIndexable) UnmarshalBinary(data []byte) (err error) {
	return proto.Unmarshal(data, (*event.Event)(e))
}

func (e *eventIndexable) Indexes() []db.Index {
	return []db.Index{
		aggregateIndex(e.AggregateID),
	}
}
func (e *eventIndexable) MasterIndexBucketKey() []byte { return []byte(eventsBucket + "MasterIndex") }

func (e aggregateIndex) BucketKey() []byte { return []byte(eventsBucket + "AggregateIndex") }
func (e aggregateIndex) Key() []byte       { return []byte(e) }

func EventKey(id uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, id)
	return b
}
