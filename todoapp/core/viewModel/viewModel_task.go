package viewModel

func (vm *ViewModel_Task) Read(r Reader) error {

	if len(vm.ID) == 0 {
		return nil
	}

	m, err := r.Task(vm.ID)
	if err != nil {
		return err
	}

	vm.ID = m.ID
	vm.Name = m.Name

	return nil
}
