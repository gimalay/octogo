package model

import "errors"

const ()

var (
	errTaskAlreadyCreated     = errors.New("Task is already created")
	errProjectAlreadyCreated  = errors.New("Project is already created")
	errActivtiyDoesNotExist   = errors.New("Task does not exist")
	errProjectDoesNotExist    = errors.New("Project does not exist")
	errAggregateID            = errors.New("cannot parse aggregate id")
	errCannotValidateEvent0   = errors.New("cannot validate event 0")
	errHashLen                = errors.New("previous massage hash unexpected len")
	errIncompatibleEvent      = errors.New("event type is incompatible with the aggregate type")
	errNoTimestamp            = errors.New("cannot store with no timestamp")
	errPrevEventHash          = errors.New("unexpected previous event hash")
	errSourceID               = errors.New("cannot parse source id")
	errSpecialPrevMessageHash = errors.New("special previous message hash expected on event 1")
	errUnexpectedID           = errors.New("unexpected ID")
	errUnknownAggregateType   = errors.New("cannot store event with unknown aggregate type")
	errUnknownType            = errors.New("cannot store event with unknown type")
	errNoAggregateFunction    = errors.New("cannot aggregate the event with unknown type")
	errEventUnknownType       = errors.New("cannot aggregate event with unknown type")
)
