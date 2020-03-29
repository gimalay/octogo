package app

import (
	db "github.com/gimalay/binx"
)

type dbMock struct {
	put    func(s db.Indexable) (err error)
	list   func(q db.QueryableSlice, cnd ...db.Condition) error
	get    func(q db.Queryable, key []byte) error
	first  func(q db.Queryable, cnd ...db.Condition) error
	last   func(q db.Queryable, cnd ...db.Condition) error
	update func(fn func(db db.Writer) error) error
	view   func(fn func(db db.Reader) error) error
}

func (m *dbMock) Put(s db.Indexable) (err error) {
	return m.put(s)
}
func (m *dbMock) List(q db.QueryableSlice, cnd ...db.Condition) error {
	return m.list(q, cnd...)
}

func (m *dbMock) Get(q db.Queryable, key []byte) error {
	return m.get(q, key)
}

func (m *dbMock) First(q db.Queryable, cnd ...db.Condition) error {
	return m.first(q, cnd...)
}

func (m *dbMock) Last(q db.Queryable, cnd ...db.Condition) error {
	return m.last(q, cnd...)
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
