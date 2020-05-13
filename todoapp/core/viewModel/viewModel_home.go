package viewModel

func (vm *ViewModel_Home) Read(r Reader) error {
	projects, err := r.Projects()
	if err != nil {
		return err
	}

	for _, p := range projects {
		if !p.Deleted {
			tasks := []*ViewModel_Home_Project_Task{}

			for _, tid := range p.Tasks {
				t, err := r.Task(tid)
				if err != nil {
					return err
				}

				tasks = append(tasks, &ViewModel_Home_Project_Task{
					ID:   t.ID,
					Name: t.Name,
				})
			}

			vm.Projects = append(vm.Projects, &ViewModel_Home_Project{
				ID:    p.ID,
				Name:  p.Name,
				Tasks: tasks,
			})
		}

	}

	return err
}
