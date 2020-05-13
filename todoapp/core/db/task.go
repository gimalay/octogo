package db

import (
	db "github.com/gimalay/binx"
	"github.com/gimalay/octogo/todoapp/core/aggregate"
	"github.com/golang/protobuf/proto"
)

type (
	task       aggregate.Task
	tasksSlice []*aggregate.Task
)

func (a *task) AppendBinary(data []byte) (bool, error) {
	return false, a.UnmarshalBinary(data)
}

func (a *task) UniqueKey() []byte { return a.ID }

const (
	TasksBucket            = "TasksBucket"
	TasksMasterIndexBucket = "TasksMasterIndexBucket"
)

func (a *tasksSlice) BucketKey() []byte { return []byte(TasksBucket) }
func (a *tasksSlice) AppendBinary(data []byte) (bool, error) {
	*a = append(*a, &aggregate.Task{})
	return true, proto.Unmarshal(data, (*a)[len(*a)-1])
}

func (a *task) Key() []byte                    { return a.ID }
func (a *task) BucketKey() []byte              { return []byte(TasksBucket) }
func (a *task) MarshalBinary() ([]byte, error) { return proto.Marshal((*aggregate.Task)(a)) }
func (a *task) UnmarshalBinary(jdata []byte) (err error) {
	return proto.Unmarshal(jdata, (*aggregate.Task)(a))
}

func (a *task) Indexes() []db.Index {
	return []db.Index{}
}
func (a *task) MasterIndexBucketKey() []byte { return []byte(TasksMasterIndexBucket) }
