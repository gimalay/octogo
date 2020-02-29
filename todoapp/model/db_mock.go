package model

import (
	"time"

	db "github.com/gimalay/binx"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/uuid"
)

type dbMock struct {
	putCallsCount int
	put           func(s db.Indexable) (err error)
	list          func(q db.QueryableSlice, cnd ...db.Condition) error
	get           func(q db.Queryable, key []byte) error
	last          func(q db.Queryable) error
	first         func(q db.Queryable) error
	update        func(fn func(db db.Writer) error) error
	view          func(fn func(db db.Reader) error) error
}

func (m *dbMock) Put(s db.Indexable) (err error) {
	m.putCallsCount++
	return m.put(s)
}
func (m *dbMock) List(q db.QueryableSlice, cnd ...db.Condition) error {
	return m.list(q, cnd...)
}

func (m *dbMock) Get(q db.Queryable, key []byte) error {
	return m.get(q, key)
}

func (m *dbMock) First(q db.Queryable, cnd ...db.Condition) error {
	return m.first(q)
}

func (m *dbMock) Last(q db.Queryable, cnd ...db.Condition) error {
	return m.last(q)
}

func (m *dbMock) Update(fn func(db db.Writer) error) error {
	return fn(m)
}
func (m *dbMock) View(fn func(db db.Reader) error) error {
	return fn(m)
}

func (m dbMock) Close() error {
	return nil
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
