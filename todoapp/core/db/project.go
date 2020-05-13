package db

import (
	db "github.com/gimalay/binx"
	"github.com/gimalay/octogo/todoapp/core/aggregate"
	"github.com/golang/protobuf/proto"
)

type (
	project       aggregate.Project
	projectsSlice []*aggregate.Project
)

func (i *project) UniqueKey() []byte { return i.ID }
func (a *project) AppendBinary(data []byte) (bool, error) {
	return false, a.UnmarshalBinary(data)
}

const (
	ProjectsBucket            = "ProjectsBucket"
	ProjectsMasterIndexBucket = "ProjectsMasterIndexBucket"
)

func (i *projectsSlice) BucketKey() []byte { return []byte(ProjectsBucket) }
func (i *projectsSlice) AppendBinary(data []byte) (bool, error) {
	*i = append(*i, &aggregate.Project{})
	return true, proto.Unmarshal(data, (*i)[len(*i)-1])
}

func (i *project) Key() []byte                    { return i.ID }
func (i *project) BucketKey() []byte              { return []byte(ProjectsBucket) }
func (i *project) MarshalBinary() ([]byte, error) { return proto.Marshal((*aggregate.Project)(i)) }
func (i *project) UnmarshalBinary(jdata []byte) (err error) {
	return proto.Unmarshal(jdata, (*aggregate.Project)(i))
}

func (i *project) Indexes() []db.Index {
	return []db.Index{}
}

func (i *project) MasterIndexBucketKey() []byte { return []byte(ProjectsMasterIndexBucket) }
