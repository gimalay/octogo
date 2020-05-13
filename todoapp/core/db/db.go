package db

import (
	"errors"
	bolt "github.com/coreos/bbolt"
	"github.com/gimalay/binx"
	"github.com/gimalay/octogo/pkg/octogo"
	"github.com/gimalay/octogo/todoapp/core/aggregate"
)

type DB struct {
	binx.Reader
	binx.Writer
}

func (w *DB) WriteAggregate(ag octogo.Aggregate) error {
	switch a := ag.(type) {
	case *aggregate.Project:
		return w.Put((*project)(a))
	case *aggregate.Task:
		return w.Put((*task)(a))
	}

	return errors.New("cannot store aggregate")
}
func (r *DB) Task(id []byte) (*aggregate.Task, error) {
	a := &task{}
	err := r.Get(a, id)
	return (*aggregate.Task)(a), err
}

func (r *DB) Project(id []byte) (*aggregate.Project, error) {
	a := &project{}
	err := r.Get(a, id)
	return (*aggregate.Project)(a), err
}
func (r *DB) ProjectByName(name string) (*aggregate.Project, error) {
	identities, err := r.Projects()
	if err != nil {
		return nil, err
	}

	for _, i := range identities {
		if i.Name == name {
			return i, nil
		}
	}

	return nil, nil
}

func (r *DB) Tasks() ([]*aggregate.Task, error) {
	a := tasksSlice{}
	err := r.Scan(&a, nil)
	return a, err
}

func (r *DB) Projects() ([]*aggregate.Project, error) {
	a := projectsSlice{}
	err := r.Scan(&a, nil)
	return a, err
}

func Schema() [][]byte {
	structure := []binx.Indexable{
		&task{},
		&project{},
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

func Open(fileName string, structure []binx.Indexable) (*bolt.DB, error) {
	buckets := [][]byte{}
	for _, b := range structure {

		buckets = append(buckets, b.BucketKey())
		buckets = append(buckets, b.MasterIndexBucketKey())
		for _, i := range b.Indexes() {
			buckets = append(buckets, i.BucketKey())
		}
	}

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
