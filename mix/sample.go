package mix

import (
	"encoding/binary"
	"math"
)

var sampleOutSpec *AudioSpec
var sampleOutNextCallback SampleOutNextCallbackFunc

type Value float64
type Sample struct {
	Values []Value
}

type SampleOutNextCallbackFunc func() []Value

func SampleConfigureOutput(s AudioSpec) {
	sampleOutSpec = &s
}

func SampleNew(values []Value) Sample {
	return Sample{
		Values: values,
	}
}

func (this Value) Abs() Value {
	return Value(math.Abs(float64(this)))
}

func (this Value) ToByteU8() byte {
	return byte(this.ToUint8())
}

func (this Value) ToByteS8() byte {
	return byte(this.ToInt8())
}

func (this Value) ToBytesU16LSB() (out []byte) {
	out = make([]byte, 2)
	binary.LittleEndian.PutUint16(out, this.ToUint16())
	return
}

func (this Value) ToBytesS16LSB() (out []byte) {
	out = make([]byte, 2)
	binary.LittleEndian.PutUint16(out, uint16(this.ToInt16()))
	return
}

func (this Value) ToBytesS32LSB() (out []byte) {
	out = make([]byte, 4)
	binary.LittleEndian.PutUint32(out, uint32(this.ToInt32()))
	return
}

func (this Value) ToBytesF32LSB() (out []byte) {
	out = make([]byte, 4)
	binary.LittleEndian.PutUint32(out, math.Float32bits(float32(this)))
	return
}

func (this Value) ToBytesF64LSB() (out []byte) {
	out = make([]byte, 4)
	binary.LittleEndian.PutUint64(out, math.Float64bits(float64(this)))
	return
}

func (this Value) ToUint8() uint8 {
	return uint8(0x80 * (this + 1))
}

func (this Value) ToInt8() int8 {
	return int8(0x80 * this)
}

func (this Value) ToUint16() uint16 {
	return uint16(0x8000 * (this + 1))
}

func (this Value) ToInt16() int16 {
	return int16(0x8000 * this)
}

func (this Value) ToInt32() int32 {
	return int32(0x80000000 * this)
}

func ValueOfByteU8(sample byte) Value {
	return Value(int8(sample))/Value(0x7F) - Value(1)
}

func ValueOfByteS8(sample byte) Value {
	return Value(int8(sample)) / Value(0x7F)
}

func ValueOfBytesU16LSB(sample []byte) Value {
	return Value(binary.LittleEndian.Uint16(sample))/Value(0x8000) - Value(1)
}

func ValueOfBytesS16LSB(sample []byte) Value {
	return Value(int16(binary.LittleEndian.Uint16(sample))) / Value(0x7FFF)
}

func ValueOfBytesS32LSB(sample []byte) Value {
	return Value(int32(binary.LittleEndian.Uint32(sample))) / Value(0x7FFFFFFF)
}

func ValueOfBytesF32LSB(sample []byte) Value {
	return Value(math.Float32frombits(binary.LittleEndian.Uint32(sample)))
}

func ValueOfBytesF64LSB(sample []byte) Value {
	return Value(math.Float64frombits(binary.LittleEndian.Uint64(sample)))
}

func SampleSetOutputCallback(fn SampleOutNextCallbackFunc) {
	sampleOutNextCallback = fn
}

func SampleOutNext() []Value {
	return sampleOutNextCallback()
}

func SampleOutNextBytes() (out []byte) {
	in := sampleOutNextCallback()
	for ch := 0; ch < sampleOutSpec.Channels; ch++ {
		switch sampleOutSpec.Format {
		case AudioU8:
			out = append(out, in[ch].ToByteU8())
		case AudioS8:
			out = append(out, in[ch].ToByteS8())
		case AudioS16:
			out = append(out, in[ch].ToBytesS16LSB()...)
		case AudioU16:
			out = append(out, in[ch].ToBytesU16LSB()...)
		case AudioS32:
			out = append(out, in[ch].ToBytesS32LSB()...)
		case AudioF32:
			out = append(out, in[ch].ToBytesF32LSB()...)
		case AudioF64:
			out = append(out, in[ch].ToBytesF64LSB()...)
		}
	}
	return
}
