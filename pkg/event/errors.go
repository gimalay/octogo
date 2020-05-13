package event

import "errors"

var (
	errAggregateID            = errors.New("cannot parse aggregate id")
	errHashLen                = errors.New("previous massage hash unexpected len")
	errNoTimestamp            = errors.New("cannot store event with no timestamp")
	errSourceID               = errors.New("cannot parse source id")
	errSpecialPrevMessageHash = errors.New("special previous message hash expected on event 1")
)
