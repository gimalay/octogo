package binding

import (
	db "github.com/gimalay/binx"
	"github.com/gimalay/octogo/todoapp/model"
)

func (vm *ViewModel_Project) read(s *Settings, r db.Reader) error {

	if len(vm.ID) == 0 {
		return nil
	}

	m := &model.Project{}

	err := r.Get(m, vm.ID)
	if err != nil {
		return err
	}

	vm.Name = m.Name

	for _, aid := range m.Tasks {

		a := &model.Task{}
		err := r.Get(a, aid)
		if err != nil {
			return err
		}
		ma, err := mapProjectTask(a, r)
		if err != nil {
			return err
		}

		vm.Tasks = append(vm.Tasks, ma)
	}

	return nil
}

func mapProjectTask(a *model.Task, r db.Reader) (*ViewModel_Project_Task, error) {
	return &ViewModel_Project_Task{
		ID:    a.ID,
		Name:  a.Name,
		Emoji: a.Emoji,
	}, nil
}
