package app

import (
	db "github.com/gimalay/binx"
	"github.com/gimalay/octogo/todoapp/model"
	"github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type (
	Binding interface {
		ViewModel(location []byte) ([]byte, error)
		Execute(command []byte) error
		Close() error
	}

	command interface {
		proto.Message
		Execute(ctx commandContext, r db.Reader) ([]newEvent, error)
	}

	binding struct {
		ctx      bindingContext
		settings *Settings
		db       db.DB
	}

	bindingContext struct {
		sourceID []byte
	}

	newEvent struct {
		Type          model.EventType
		Payload       proto.Message
		AggregateID   []byte
		AggregateType model.AggregateType
	}

	commandContext struct {
	}

	eventContext struct {
		sourceID  []byte
		timestamp *timestamp.Timestamp
	}

	viewModel interface {
		read(*Settings, db.Reader) error
		proto.Message
	}
)

func New(fileName string) (Binding, error) {
	db, err := openDB(fileName)
	if err != nil {
		return nil, err
	}

	ctx := bindingContext{
		sourceID: sourceId,
	}

	bnd := &binding{
		settings: &Settings{},
		ctx:      ctx,
		db:       db,
	}

	return bnd, err
}

func (b *binding) Execute(data []byte) error {
	cmd := &Command{}
	err := proto.Unmarshal(data, cmd)
	if err != nil {
		return err
	}

	if cmd.Timestamp == nil {
		return errors.New("not timestamp on command")
	}

	ex := unmarshal(cmd)
	err = b.db.Update(func(w db.Writer) error {
		err = execute(b.ctx, cmd.Timestamp, w, ex)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return errors.Wrap(err, "command exeution")
	}

	return nil
}

func (s *binding) ViewModel(location []byte) ([]byte, error) {

	loc := &Location{}
	err := proto.Unmarshal(location, loc)
	if err != nil {
		panic("cannot unmarshal location")
	}

	vm := loc.mapToViewModel()

	err = s.db.View(func(r db.Reader) (err error) {
		return vm.read(s.settings, r)
	})

	if err != nil {
		return []byte{}, err
	}

	return proto.Marshal(vm)
}

func openDB(fileName string) (db.DB, error) {
	return db.Open(fileName, []db.Indexable{
		&model.Event{},
		&model.Task{},
		&model.Project{},
	})
}

func (b *binding) Close() error {
	return b.db.Close()
}

func id(b byte) []byte {
	v, _ := uuid.NewMD5(uuid.NameSpaceDNS, []byte{b}).MarshalBinary()
	return v
}

var sourceId = id(1)
