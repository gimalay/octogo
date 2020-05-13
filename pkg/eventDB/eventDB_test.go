package eventDB

import (
	"os"
	"testing"
	"time"

	bolt "github.com/coreos/bbolt"
	"github.com/gimalay/binx"
	"github.com/gimalay/octogo/pkg/aggregator"
	"github.com/gimalay/octogo/pkg/event"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

// {
// 		"event 1 has zeros previous hash",
// 		&Event{
// 			ID:                  1,
// 			PreviousMessageHash: make([]byte, 32),
// 			SourceID:            id(1),
// 			AggregateID:         id(3),
// 			Type:                1,
// 			Timestamp:           &timestamp.Timestamp{},
// 		},
// 		errSpecialPrevMessageHash,
// 	},
func Test_appendBlock(t *testing.T) {
	tests := []struct {
		name     string
		existing *eventIndexable
		argument *event.Event
		err      error
	}{
		{
			name:     "accept nil as prev event",
			existing: nil,
			argument: &event.Event{
				ID:                  1,
				PreviousMessageHash: make([]byte, 32),
				AggregateID:         id(1),
				SourceID:            id(1),
				Timestamp:           tm(1),
			},
			err: nil,
		},
		{
			name: "append event if it is valid and has a matching hash and ID is previous ID + 1",
			existing: &eventIndexable{
				ID:          1,
				AggregateID: id(1),
			},
			argument: &event.Event{
				ID: 2,
				PreviousMessageHash: eventHash(&event.Event{
					ID:          1,
					AggregateID: id(1),
				}),
				AggregateID: id(1),
				SourceID:    id(1),
				Type:        1,
				Timestamp:   tm(1),
			},
			err: nil,
		},
		{
			name: "return error is ID is not squential",
			existing: &eventIndexable{
				ID:          1,
				AggregateID: id(1),
			},
			argument: &event.Event{
				ID: 3,
				PreviousMessageHash: eventHash(&event.Event{
					ID:          1,
					AggregateID: id(1),
				}),
				AggregateID: id(1),
				SourceID:    id(1),
				Type:        1,
				Timestamp:   tm(1),
			},
			err: errUnexpectedID,
		},

		{
			name: "append event with no hash for previous event",
			existing: &eventIndexable{
				ID:          2,
				AggregateID: id(1),
			},
			argument: &event.Event{
				ID:                  3,
				PreviousMessageHash: make([]byte, 32),
				AggregateID:         id(1),
				SourceID:            id(1),
				Type:                1,
				Timestamp:           tm(1),
			},
			err: errPrevEventHash,
		},
		{
			name: "return error if previous event hash is diffrent",
			existing: &eventIndexable{
				ID:          2,
				AggregateID: id(1),
			},
			argument: &event.Event{
				ID:                  3,
				PreviousMessageHash: eventHash(&event.Event{ID: 4}),
				AggregateID:         id(1),
				SourceID:            id(1),
				Type:                1,
				Timestamp:           tm(1),
			},
			err: errPrevEventHash,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, teardown := prep(t)
			defer teardown()

			if tt.existing != nil {
				err := db.Update(func(tx *bolt.Tx) error {
					w := binx.NewWriter(tx)
					err := w.Put(tt.existing)
					assert.Nil(t, err)
					return nil
				})
				assert.Nil(t, err)
			}

			_ = db.Update(func(tx *bolt.Tx) error {
				err := (&EventDB{tx}).AppendEventBlock(tt.argument)
				assert.Equal(t, tt.err, err)
				return err
			})
		})
	}
}
func Test_eventsForAggregate(t *testing.T) {
	tests := []struct {
		name     string
		existing []*eventIndexable
		argument []byte
		expected []aggregator.Event
		err      error
	}{
		{
			name: "order and combine existing events by timestamp",
			existing: []*eventIndexable{
				{
					ID:          1,
					AggregateID: id(1),
					Timestamp:   tm(2),
					Payload:     []byte{2},
				},
				{
					ID:          2,
					AggregateID: id(1),
					Timestamp:   tm(1),
					Payload:     []byte{1},
				}},
			argument: id(1),
			expected: []aggregator.Event{
				{
					AggregateID: id(1),
					Timestamp:   tm(1),
					Payload:     []byte{1},
				},
				{
					AggregateID: id(1),
					Timestamp:   tm(2),
					Payload:     []byte{2},
				},
			},
		},
		{
			name: "order and combine existing with same date by ID",
			existing: []*eventIndexable{
				{
					ID:          2,
					AggregateID: id(1),
					Timestamp:   tm(1),
					Payload:     []byte{2},
				},
				{
					ID:          1,
					AggregateID: id(1),
					Timestamp:   tm(1),
					Payload:     []byte{1},
				}},
			argument: id(1),
			expected: []aggregator.Event{
				{
					AggregateID: id(1),
					Timestamp:   tm(1),
					Payload:     []byte{1},
				},
				{
					AggregateID: id(1),
					Timestamp:   tm(1),
					Payload:     []byte{2},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, teardown := prep(t)
			defer teardown()

			err := db.Update(func(tx *bolt.Tx) error {
				w := binx.NewWriter(tx)
				for _, e := range tt.existing {
					err := w.Put(e)
					assert.Nil(t, err)
				}

				return nil
			})
			assert.Nil(t, err)

			err = db.View(func(tx *bolt.Tx) error {
				es, err := (&EventDB{tx}).EventsForAggregate(tt.argument)
				assert.Nil(t, err)
				assert.Equal(t, tt.expected, es)
				return err
			})
		})
	}
}

func eventHash(e *event.Event) []byte {
	hash, _ := e.Hash()
	return hash[:]
}

const dbPath = "/tmp/testdb"

func prep(t *testing.T) (*bolt.DB, func()) {
	_ = os.Remove(dbPath)
	b, err := open(dbPath, Schema())
	assert.Nil(t, err)

	return b, func() {
		err = b.Close()
		assert.Nil(t, err)
	}
}
func open(fileName string, buckets [][]byte) (*bolt.DB, error) {
	bdb, err := bolt.Open(fileName, 0600, nil)
	s := bdb
	if err != nil {
		return nil, err
	}

	err = s.Update(func(tx *bolt.Tx) (err error) {
		for _, v := range buckets {
			if _, err = tx.CreateBucketIfNotExists(v); err != nil {
				return err
			}
		}

		return err
	})

	return bdb, nil
}
func tm(s int64) time.Time {
	return time.Time{}.Add(time.Duration(s))
}

func ttm(s int64) time.Time {
	return time.Time{}.Add(time.Second * time.Duration(s))
}

func bt(message proto.Message) []byte {
	b, _ := proto.Marshal(message)
	return b
}
