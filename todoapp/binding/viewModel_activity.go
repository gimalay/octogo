package binding

import (
	db "github.com/gimalay/binx"
	"github.com/gimalay/octogo/todoapp/model"
)

func (vm *ViewModel_Task) read(s *Settings, r db.Reader) error {

	if len(vm.ID) == 0 {
		return nil
	}

	m := &model.Task{}

	err := r.Get(m, vm.ID)
	if err != nil {
		return err
	}

	vm.ID = m.ID
	vm.Name = m.Name
	vm.Emoji = m.Emoji

	return nil
}
