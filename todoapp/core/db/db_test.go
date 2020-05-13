package db

import (
	"os"
	"testing"
	"time"

	bolt "github.com/coreos/bbolt"
	"github.com/gimalay/binx"
	"github.com/gimalay/insight/internal/aggregate"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func ttm(t time.Time) *timestamp.Timestamp {
	r, _ := ptypes.TimestampProto(t)
	return r
}

func tm(t *timestamp.Timestamp) time.Time {
	r, _ := ptypes.Timestamp(t)
	return r
}

func Test_EntriesToActivities(t *testing.T) {
	tests := []struct {
		name     string
		existing []binx.Indexable
		expected []*aggregate.Activity
	}{
		{
			name: "no activity",
			existing: []binx.Indexable{
				&entry{
					ID:         []byte{2},
					ActivityID: []byte{1},
					Start:      time.Time{}.Add(1),
				},
			},
			expected: activitiesSlice{nil},
		},
		{
			name: "join by ID",
			existing: []binx.Indexable{
				&activity{
					ID: []byte{1},
				},
				&entry{
					ID:         []byte{2},
					ActivityID: []byte{1},
					Start:      time.Time{}.Add(1),
				},
			},
			expected: activitiesSlice{
				&aggregate.Activity{
					ID: []byte{1},
				},
			},
		},
		{
			name: "join by ID first",
			existing: []binx.Indexable{
				&activity{
					ID: []byte{1},
				},
				&activity{
					ID:                 []byte{2},
					PrimaryWorkoutType: 1,
				},
				&entry{
					ID:          []byte{2},
					ActivityID:  []byte{1},
					WorkoutType: 1,
					Start:       time.Time{}.Add(1),
				},
			},
			expected: activitiesSlice{
				&aggregate.Activity{
					ID: []byte{1},
				},
			},
		},
		{
			name: "join by PrimaryWorkutType",
			existing: []binx.Indexable{
				&activity{
					ID:                 []byte{1},
					PrimaryWorkoutType: 1,
				},
				&entry{
					ID:          []byte{2},
					WorkoutType: 1,
					Start:       time.Time{}.Add(1),
				},
			},
			expected: activitiesSlice{
				&aggregate.Activity{
					ID:                 []byte{1},
					PrimaryWorkoutType: 1,
				},
			},
		},
		{
			name: "join by SecondaryWorkutType",
			existing: []binx.Indexable{
				&activity{
					ID:                    []byte{1},
					SecondaryWorkoutTypes: []int32{2, 1},
				},
				&entry{
					ID:          []byte{2},
					WorkoutType: 1,
					Start:       time.Time{}.Add(1),
				},
			},
			expected: []*aggregate.Activity{
				&aggregate.Activity{
					ID:                    []byte{1},
					SecondaryWorkoutTypes: []int32{2, 1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, teardown := prep(t)
			defer teardown()

			err := db.Update(func(tx *bolt.Tx) error {
				for _, e := range tt.existing {
					w := binx.NewWriter(tx)
					err := w.Put(e)
					assert.Nil(t, err)
				}

				return nil
			})
			assert.Nil(t, err)

			err = db.View(func(tx *bolt.Tx) error {
				r := binx.NewReader(tx)
				w := binx.NewWriter(tx)
				es, err := (&DB{r, w}).Entries(1000)
				assert.Nil(t, err)
				as, err := (&DB{r, w}).EntriesToActivities(es)
				assert.Nil(t, err)
				assert.Equal(t, tt.expected, as)
				return err
			})
			assert.Nil(t, err)

		})
	}
}

const (
	start int64 = 25340230080
)

func Test_CountActivityEntriesInRange(t *testing.T) {
	tests := []struct {
		name     string
		existing []binx.Indexable
		from, to time.Time
		expected int
	}{
		{
			name: "Middle range",
			existing: []binx.Indexable{
				&activity{
					ID: id(1),
				},
				&entry{
					ID:         []byte{1},
					ActivityID: id(1),
					Start:      time.Time{}.Add(1 * time.Second),
				},
				&entry{
					ID:         []byte{2},
					ActivityID: id(1),
					Start:      time.Time{}.Add(2 * time.Second),
				},
				&entry{
					ID:         []byte{3},
					ActivityID: id(1),
					Start:      time.Time{}.Add(3 * time.Second),
				},
				&entry{
					ID:         []byte{4},
					ActivityID: id(1),
					Start:      time.Time{}.Add(4 * time.Second),
				},
			},
			expected: 2,
			from:     time.Time{}.Add(2 * time.Second),
			to:       time.Time{}.Add(3 * time.Second),
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
				r := binx.NewReader(tx)
				w := binx.NewWriter(tx)
				es, err := (&DB{r, w}).CountActivityEntriesInRange(id(1), tt.from, tt.to)
				assert.Nil(t, err)
				assert.Equal(t, tt.expected, es)
				return err
			})
			assert.Nil(t, err)

		})
	}
}

func Test_EntriesRange(t *testing.T) {
	tests := []struct {
		name     string
		existing []binx.Indexable
		from, to time.Time
		expected []*aggregate.Entry
	}{
		{
			name: "Middle range",
			existing: []binx.Indexable{
				&entry{
					ID:    []byte{1},
					Start: time.Time{}.Add(1 * time.Second),
				},
				&entry{
					ID:    []byte{2},
					Start: time.Time{}.Add(2 * time.Second),
				},
				&entry{
					ID:    []byte{3},
					Start: time.Time{}.Add(3 * time.Second),
				},
				&entry{
					ID:    []byte{4},
					Start: time.Time{}.Add(4 * time.Second),
				},
			},
			expected: []*aggregate.Entry{
				{
					ID:    []byte{2},
					Start: time.Time{}.Add(2 * time.Second),
				},
				{
					ID:    []byte{3},
					Start: time.Time{}.Add(3 * time.Second),
				},
			},
			from: time.Time{}.Add(2 * time.Second),
			to:   time.Time{}.Add(3 * time.Second),
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
				r := binx.NewReader(tx)
				w := binx.NewWriter(tx)
				es, err := (&DB{r, w}).EntriesRange(tt.from, tt.to)
				assert.Nil(t, err)
				assert.Equal(t, tt.expected, es)
				return err
			})
			assert.Nil(t, err)

		})
	}
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
func id(b byte) []byte {
	v, _ := uuid.NewMD5(uuid.NameSpaceDNS, []byte{b}).MarshalBinary()
	return v

}
