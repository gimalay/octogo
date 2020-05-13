package viewModel

func (vm *ViewModel_Project) Read(r Reader) error {

	if len(vm.ID) == 0 {
		return nil
	}

	m, err := r.Project(vm.ID)
	if err != nil {
		return err
	}

	vm.Name = m.Name

	for _, aid := range m.Tasks {

		ma, err := r.Task(aid)
		if err != nil {
			return err
		}

		vm.Tasks = append(vm.Tasks, &ViewModel_Project_Task{
			ID:   ma.ID,
			Name: ma.Name,
		})
	}

	return nil
}
