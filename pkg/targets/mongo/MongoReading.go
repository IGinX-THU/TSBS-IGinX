// automatically generated by the FlatBuffers compiler, do not modify

package mongo

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type MongoReading struct {
	_tab flatbuffers.Table
}

func GetRootAsMongoReading(buf []byte, offset flatbuffers.UOffsetT) *MongoReading {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MongoReading{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *MongoReading) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MongoReading) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *MongoReading) Key() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *MongoReading) Value() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *MongoReading) MutateValue(n float64) bool {
	return rcv._tab.MutateFloat64Slot(6, n)
}

func MongoReadingStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func MongoReadingAddKey(builder *flatbuffers.Builder, key flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(key), 0)
}
func MongoReadingAddValue(builder *flatbuffers.Builder, value float64) {
	builder.PrependFloat64Slot(1, value, 0.0)
}
func MongoReadingEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
