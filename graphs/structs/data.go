package structs

import "sync/atomic"

type DataTag int64

var currentTag int64 = 0

func NextTag() DataTag {
	atomic.AddInt64(&currentTag, 1)
	return DataTag(currentTag)
}

type IDataItem interface {
	Get
}

type VertexData struct {
	map[DataTag]IDataItem
}

func (d *VertexData) Get[D any](tag DataTag) (D, bool) {

}
