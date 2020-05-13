package eventDB

import "errors"

var (
	errPrevEventHash        = errors.New("unexpected previous event hash")
	errCannotValidateEvent0 = errors.New("cannot validate event 0")
	errUnexpectedID         = errors.New("unexpected ID")
)
