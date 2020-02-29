package binding

import (
	"bytes"

	db "github.com/gimalay/binx"
	"github.com/gimalay/octogo/todoapp/model"
)

func (state *ViewModel_Home) read(settings *Settings, r db.Reader) error {
	Tasks := model.TasksSlice{}
	err := r.List(&Tasks)
	if err != nil {
		return err
	}
	identites := model.ProjectsSlice{}
	err = r.List(&identites)
	if err != nil {
		return err
	}

	is, err := homeProjectsState(identites, Tasks, r)
	if err != nil {
		return err
	}

	state.Projects = is

	if err != nil {
		return err
	}
	return err
}

func homeProjectsState(identites model.ProjectsSlice, as model.TasksSlice, r db.Reader) ([]*ViewModel_Home_Project, error) {
	ar := []*ViewModel_Home_Project{}
	for _, i := range identites {
		hia := homeProjectTasks(i, as)

		his, err := homeProjectState(i, hia, r)
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

func homeProjectState(i *model.Project, as model.TasksSlice, r db.Reader) (*ViewModel_Home_Project, error) {
	ar := []*ViewModel_Home_Project_Task{}
	for _, a := range as {
		as, err := homeTaskState(a, r)
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
func homeTaskState(a *model.Task, r db.Reader) (*ViewModel_Home_Project_Task, error) {
	return &ViewModel_Home_Project_Task{
		ID:    a.ID,
		Name:  a.Name,
		Emoji: a.Emoji,
	}, nil
}
