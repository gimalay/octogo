package model

import (
	"crypto/sha256"
	"encoding/binary"
	"errors"

	db "github.com/gimalay/binx"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
)

const Version int32 = 1

type (
	EventSlice          []Event
	EventAggregateIndex []byte
)

type Hash [sha256.Size]byte

const (
	EventsBucket              = "Events"
	EventAggregateIndexBucket = "EventAggregateIndexBucket"
	EventsMasterIndexBucket   = "EventsMasterIndexBucket"
)

func (s *EventSlice) BucketKey() []byte { return []byte(EventsBucket) }
func (s *EventSlice) AppendBinary(data []byte) error {
	*s = append(*s, Event{})
	return proto.Unmarshal(data, &(*s)[len(*s)-1])
}

func (e *Event) Key() []byte {
	return EventKey(e.ID)
}
func (e *Event) BucketKey() []byte { return []byte(EventsBucket) }
func (e *Event) MarshalBinary() ([]byte, error) {
	return proto.Marshal(e)
}
func (e *Event) UnmarshalBinary(jdata []byte) (err error) { return proto.Unmarshal(jdata, e) }

func (e *Event) Indexes() []db.Index {
	return []db.Index{
		EventAggregateIndex(e.AggregateID),
	}
}
func (e *Event) MasterIndexBucketKey() []byte { return []byte(EventsMasterIndexBucket) }

func (e EventAggregateIndex) BucketKey() []byte { return []byte(EventAggregateIndexBucket) }
func (e EventAggregateIndex) Key() []byte       { return []byte(e) }

func (e *Event) ValidateChain(prev *Event) error {
	if err := e.Validate(); err != nil {
		return err
	}

	if e.ID != prev.ID+1 {
		return errUnexpectedID
	}

	hash, err := prev.Hash()

	if err != nil {
		return err
	}
	if e.PrevHash() != hash {
		return errPrevEventHash
	}
	return nil
}

func validateEventType(et EventType, at AggregateType) bool {
	switch at {
	case AggregateType_Project:
		return mapProjectEvent(et) != nil
	case AggregateType_Task:
		return mapTaskEvent(et) != nil
	}
	return false
}

func (e *Event) Validate() error {
	if e.ID == 0 {
		return errors.New("cannot validate event 0")
	}

	emptyHash := Hash{}

	if e.ID == 1 && e.PrevHash() != emptyHash {
		return errSpecialPrevMessageHash
	}
	_, err := uuid.FromBytes(e.SourceID)
	if err != nil {
		return errSourceID
	}
	_, err = uuid.FromBytes(e.AggregateID)
	if err != nil {
		return errAggregateID
	}
	if len(e.PreviousMessageHash) != 32 {
		return errHashLen
	}
	if e.Type == EventType_Unknown && e.SourceVersion <= Version {
		return errUnknownType
	}
	if e.AggregateType == AggregateType_UnknownAggregate && e.SourceVersion <= Version {
		return errUnknownAggregateType
	}
	if !validateEventType(e.Type, e.AggregateType) && e.SourceVersion <= Version {
		return errIncompatibleEvent
	}
	if e.Timestamp == nil {
		return errNoTimestamp
	}
	//TODO: Validate payload
	return nil
}

func (e *Event) Hash() (Hash, error) {
	if e.ID == 0 {
		//special hash for event zero
		return Hash{}, nil
	}
	data, err := proto.Marshal(e)
	if err != nil {
		return Hash{}, err
	}

	return sha256.Sum256(data), nil
}

func (e *Event) PrevHash() Hash {
	hash := Hash{}
	copy(hash[:], e.PreviousMessageHash)
	return hash
}

func EventKey(id uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, id)
	return b
}

func AppendEvent(writer db.Writer, event, prev *Event) error {
	if err := event.ValidateChain(prev); err != nil {
		return err
	}
	return writer.Put(event)
}
