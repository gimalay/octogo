package binding

import (
	"github.com/gimalay/octogo/todoapp/app"
)

type (
	Binding interface {
		ViewModel(location []byte) ([]byte, error)
		Execute(command []byte) error
		Close() error
	}
)

func New(fileName string) (Binding, error) {
	return app.New(fileName)
}

