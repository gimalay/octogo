package event

import (
	"crypto/sha256"
	"errors"
	time "time"

	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
)

type Hash [sha256.Size]byte

func (e *Event) Validate() error {
	if e.ID == 0 {
		return errors.New("cannot validate event 0")
	}

	emptyHash := Hash{}

	if e.ID == 1 && e.PrevHash() != emptyHash {
		return errSpecialPrevMessageHash
	}
	_, err := uuid.FromBytes(e.SourceID)
	if err != nil {
		return errSourceID
	}
	_, err = uuid.FromBytes(e.AggregateID)
	if err != nil {
		return errAggregateID
	}
	if len(e.PreviousMessageHash) != 32 {
		return errHashLen
	}
	if e.Timestamp.Equal(time.Time{}) {
		return errNoTimestamp
	}
	return nil
}

func (e *Event) Hash() (Hash, error) {
	if e.ID == 0 {
		//special hash for event zero
		return Hash{}, nil
	}
	data, err := proto.Marshal(e)
	if err != nil {
		return Hash{}, err
	}

	return sha256.Sum256(data), nil
}

func (e *Event) PrevHash() Hash {
	hash := Hash{}
	copy(hash[:], e.PreviousMessageHash)
	return hash
}
