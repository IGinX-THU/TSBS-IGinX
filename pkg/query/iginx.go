package query

import (
	"fmt"
	"sync"
)

// Iginx encodes a Iginx request. This will be serialized for use
// by the tsbs_run_queries_iginx program.
type Iginx struct {
	HumanLabel       []byte
	HumanDescription []byte

	SqlQuery   []byte
	id         uint64
}

// IginxPool is a sync.Pool of Iginx Query types
var IginxPool = sync.Pool{
	New: func() interface{} {
		return &Iginx{
			HumanLabel:       make([]byte, 0, 1024),
			HumanDescription: make([]byte, 0, 1024),
			SqlQuery:         make([]byte, 0, 1024),
		}
	},
}

// NewIginx returns a new Iginx Query instance
func NewIginx() *Iginx {
	return IginxPool.Get().(*Iginx)
}

// GetID returns the ID of this Query
func (q *Iginx) GetID() uint64 {
	return q.id
}

// SetID sets the ID for this Query
func (q *Iginx) SetID(n uint64) {
	q.id = n
}

// String produces a debug-ready description of a Query.
func (q *Iginx) String() string {
	return fmt.Sprintf("HumanLabel: %s, HumanDescription: %s, Query: %s", q.HumanLabel, q.HumanDescription, q.SqlQuery)
}

// HumanLabelName returns the human readable name of this Query
func (q *Iginx) HumanLabelName() []byte {
	return q.HumanLabel
}

// HumanDescriptionName returns the human readable description of this Query
func (q *Iginx) HumanDescriptionName() []byte {
	return q.HumanDescription
}

// Release resets and returns this Query to its pool
func (q *Iginx) Release() {
	q.HumanLabel = q.HumanLabel[:0]
	q.HumanDescription = q.HumanDescription[:0]
	q.id = 0
	q.SqlQuery = q.SqlQuery[:0]

	IginxPool.Put(q)
}

