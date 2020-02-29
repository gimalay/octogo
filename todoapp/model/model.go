package model

import (
	"github.com/google/uuid"
)

func NewID() []byte {
	id := uuid.New()
	v, err := id.MarshalBinary()
	if err != nil {
		panic("cannot generate uuid")
	}
	return v
}
