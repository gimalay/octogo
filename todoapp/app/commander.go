package app

import (
	db "github.com/gimalay/binx"
	"github.com/gimalay/octogo/todoapp/model"
	"github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

func unmarshal(cmd *Command) command {
	c := mapCommandPayload(cmd.Type)
	err := proto.Unmarshal(cmd.Payload, c)
	if err != nil {
		panic("cannot unmarshal")
	}
	return c
}

func execute(ctx bindingContext, t *timestamp.Timestamp, w db.Writer, cmd command) error {
	events, err := cmd.Execute(commandContext{}, w)

	if err != nil {
		return err
	}

	return appendEvents(eventContext{sourceID: ctx.sourceID, timestamp: t}, w, events)
}

func appendEvents(ctx eventContext, w db.Writer, events []newEvent) error {
	for _, e := range events {
		err := appendEvent(ctx, w, e)
		if err != nil {
			return err
		}
		err = model.Aggregate(w, e.AggregateID)
		if err != nil {
			return err
		}
	}
	return nil
}

func appendEvent(ctx eventContext, w db.Writer, ev newEvent) error {
	prev := &model.Event{} //event zero has all default values
	lastHash := model.Hash{}
	err := w.Last(prev)
	if err != nil && err != db.ErrNotFound {
		return err
	}

	if err != db.ErrNotFound {
		lastHash, err = prev.Hash()
		if err != nil {
			return err
		}
	}

	data, err := proto.Marshal(ev.Payload)
	if err != nil {
		return err
	}

	m := &model.Event{
		SourceID:            ctx.sourceID,
		SourceVersion:       model.Version,
		AggregateID:         ev.AggregateID,
		PreviousMessageHash: lastHash[:],
		Timestamp:           ctx.timestamp,
		ID:                  prev.ID + 1,
		Type:                ev.Type,
		Payload:             data,
		AggregateType:       ev.AggregateType,
	}

	return model.AppendEvent(w, m, prev)
}
