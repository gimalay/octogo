package aggregate

import "errors"

var (
	errTaskAlreadyCreated    = errors.New("Task is already created")
	errProjectAlreadyCreated = errors.New("Project is already created")
	errTaskDoesNotExist      = errors.New("Task does not exist")
	errProjectDoesNotExist   = errors.New("Project does not exist")

	errEventUnknownType = errors.New("cannot aggregate event with unknown type")
)
