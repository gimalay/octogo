package app

import (
	"fmt"
	"github.com/gimalay/octogo/pkg/aggregator"
	"github.com/gimalay/octogo/pkg/eventDB"
	"github.com/gimalay/octogo/todoapp/core/aggregate"
	"github.com/gimalay/octogo/todoapp/core/command"
	"github.com/gimalay/octogo/todoapp/core/db"
	"github.com/gimalay/octogo/todoapp/core/viewModel"
	"time"

	bolt "github.com/coreos/bbolt"
	"github.com/gimalay/binx"
)

type Message struct {
	Type      int
	Payload   []byte
	Timestamp time.Time
}

type App struct {
	db *bolt.DB
}

func (a *App) ViewModel(location Message) ([]byte, error) {
	var vm []byte

	err := a.db.View(func(tx *bolt.Tx) (err error) {
		reader := &db.DB{Reader: binx.NewReader(tx), Writer: binx.NewWriter(tx)}

		nav := viewModel.Navigator{
			Reader: reader,
		}
		vm, err = nav.NavigateTo(viewModel.Message(location))
		return err
	})

	if err != nil {
		fmt.Printf("view model error: %v \n", err.Error())
		return []byte{}, err
	}

	return vm, nil
}

func (a *App) Execute(cmd Message) error {
	return a.db.Update(func(tx *bolt.Tx) error {
		oc := aggregator.Aggregator{
			EventWriter:     &eventDB.EventDB{Tx: tx},
			EventReader:     &eventDB.EventDB{Tx: tx},
			AggregateWriter: &db.DB{Writer: binx.NewWriter(tx)},
			AggregateMapper: &aggregate.Mapper{},
		}

		c, err := (&command.Unmarshaler{Reader: &db.DB{Reader: binx.NewReader(tx)}}).Unmarshal(command.Message(cmd))
		if err != nil {
			return err
		}
		return oc.Execute(c)
	})
}

func New(fileName string) (*App, error) {
	reader, err := openDB(fileName, append(db.Schema(), eventDB.Schema()...))
	if err != nil {
		return nil, err
	}

	a := &App{
		db: reader,
	}

	return a, err
}

func (b *App) Close() error {
	return b.db.Close()
}

func openDB(fileName string, buckets [][]byte) (*bolt.DB, error) {
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
