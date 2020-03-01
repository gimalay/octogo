package binding

import (
	"bytes"

	db "github.com/gimalay/binx"
	"github.com/gimalay/octogo/todoapp/model"
)

func (vm *ViewModel_Home) read(settings *Settings, r db.Reader) error {
	tasks := model.TasksSlice{}
	err := r.List(&tasks)
	if err != nil {
		return err
	}
	projects := model.ProjectsSlice{}
	err = r.List(&projects)
	if err != nil {
		return err
	}

	activProjects := model.ProjectsSlice{}

	for _, p := range projects {
		if !p.Deleted {
			activProjects = append(activProjects, p)
		}

	}

	is, err := homeProjectsViewModel(activProjects, tasks, r)
	if err != nil {
		return err
	}

	vm.Projects = is

	if err != nil {
		return err
	}
	return err
}

func homeProjectsViewModel(projects model.ProjectsSlice, as model.TasksSlice, r db.Reader) ([]*ViewModel_Home_Project, error) {
	ar := []*ViewModel_Home_Project{}
	for _, i := range projects {
		hia := homeProjectTasks(i, as)

		his, err := homeProjectViewModel(i, hia, r)
		if err != nil {
			return nil, err
		}

		ar = append(ar, his)
	}

	return ar, nil
}

func homeProjectTasks(i *model.Project, as model.TasksSlice) model.TasksSlice {
	res := model.TasksSlice{}
	for _, id := range i.Tasks {
		for _, a := range as {
			if bytes.Equal(a.ID, id) {

				res = append(res, a)
			}
		}
	}
	return res
}

func homeProjectViewModel(i *model.Project, as model.TasksSlice, r db.Reader) (*ViewModel_Home_Project, error) {
	ar := []*ViewModel_Home_Project_Task{}
	for _, a := range as {
		as, err := homeTaskViewModel(a, r)
		if err != nil {
			return nil, err
		}

		ar = append(ar, as)
	}

	return &ViewModel_Home_Project{
		ID:    i.ID,
		Name:  i.Name,
		Tasks: ar,
	}, nil

}
func homeTaskViewModel(a *model.Task, r db.Reader) (*ViewModel_Home_Project_Task, error) {
	return &ViewModel_Home_Project_Task{
		ID:    a.ID,
		Name:  a.Name,
		Emoji: a.Emoji,
	}, nil
}
