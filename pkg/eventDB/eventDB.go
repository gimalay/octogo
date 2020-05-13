package eventDB

import (
	"errors"
	"sort"

	bolt "github.com/coreos/bbolt"
	"github.com/gimalay/binx"
	"github.com/gimalay/octogo/pkg/aggregator"
	"github.com/gimalay/octogo/pkg/event"
	"github.com/google/uuid"
)

const Version = 1

type EventDB struct {
	*bolt.Tx
}

func (s *EventDB) WriteEvent(e aggregator.Event) error {
	lastHash := event.Hash{}
	prev, err := s.last()
	if err != nil {
		return err
	}

	if prev != nil {
		lastHash, err = prev.Hash()
		if err != nil {
			return err
		}
	} else {
		prev = &event.Event{}
	}

	ev := &event.Event{
		Timestamp:   e.Timestamp,
		Payload:     e.Payload,
		AggregateID: e.AggregateID,
		Type:        int32(e.Type),
	}

	ev.PreviousMessageHash = lastHash[:]
	ev.SourceID = id(1)
	ev.SourceVersion = Version
	ev.ID = prev.ID + 1

	return s.AppendEventBlock(ev)
}

func (s *EventDB) EventsForAggregate(aggregateID []byte) ([]aggregator.Event, error) {
	var events = EventSlice{}
	err := binx.NewReader(s.Tx).Scan(&events, []binx.Bound{
		binx.Where{Index: aggregateIndex(aggregateID)}})

	sort.Slice(events, func(i, j int) bool {
		it := events[i].Timestamp
		jt := events[j].Timestamp

		return err == nil && it.Before(jt) && jt.After(it) || events[i].ID < events[j].ID
	})

	es := []aggregator.Event{}

	if err != nil {
		return es, err
	}

	for _, e := range events {

		if e.SourceVersion > Version {
			continue
		}

		es = append(es, aggregator.Event{
			Type:        int(e.Type),
			Payload:     e.Payload,
			AggregateID: e.AggregateID,
			Timestamp:   e.Timestamp,
		})
	}

	return es, err
}

func (s *EventDB) AppendEventBlock(event *event.Event) error {
	last, err := s.last()
	if err != nil {
		return err
	}

	if err := event.Validate(); err != nil {
		return err
	}

	if err := validateChain(last, event); err != nil {
		return err
	}
	return s.put(event)
}

func (s *EventDB) put(ev *event.Event) error {
	return binx.NewWriter(s.Tx).Put((*eventIndexable)(ev))
}

func validateChain(prev *event.Event, e *event.Event) error {
	if prev == nil {
		prev = &event.Event{}
	}
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

func (s *EventDB) last() (*event.Event, error) {
	bkt := s.Tx.Bucket([]byte(eventsBucket))
	if bkt == nil {
		return nil, errors.New("cannot get events bucket")
	}
	_, data := bkt.Cursor().Last()

	ev := &eventIndexable{}
	err := ev.UnmarshalBinary(data)
	return (*event.Event)(ev), err
}

func Schema() [][]byte {
	structure := []binx.Indexable{
		&eventIndexable{},
	}
	buckets := [][]byte{}
	for _, b := range structure {
		buckets = append(buckets, b.BucketKey())
		buckets = append(buckets, b.MasterIndexBucketKey())
		for _, i := range b.Indexes() {
			buckets = append(buckets, i.BucketKey())
		}
	}
	return buckets
}

func id(b byte) []byte {
	v, _ := uuid.NewMD5(uuid.NameSpaceDNS, []byte{b}).MarshalBinary()
	return v

}
