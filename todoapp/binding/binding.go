package binding

import (
	"github.com/gimalay/octogo/todoapp/core/app"
	"time"
)

type (
	Message struct {
		Type      int
		Payload   []byte
		Timestamp string
	}

	Binding interface {
		ViewModel(t int, payload []byte, timestamp string) ([]byte, error)
		Execute(t int, payload []byte, timestamp string) error
		Close() error
	}

	binding struct {
		*app.App
	}
)

func toAppMessage(t int, payload []byte, timestamp string) (app.Message, error) {
	tm, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		return app.Message{}, err
	}

	return app.Message{Type: t, Payload: payload, Timestamp: tm}, nil
}

func (b binding) ViewModel(t int, payload []byte, timestamp string) ([]byte, error) {
	msg, err := toAppMessage(t, payload, timestamp)
	if err != nil {
		return nil, err
	}

	return b.App.ViewModel(msg)
}

func (b binding) Execute(t int, payload []byte, timestamp string) error {
	msg, err := toAppMessage(t, payload, timestamp)
	if err != nil {
		return err
	}

	return b.App.Execute(msg)
}

func New(fileName string) (Binding, error) {
	a, err := app.New(fileName)
	if err != nil {
		return nil, err
	}
	bnd := binding{App: a}

	return bnd, err
}
