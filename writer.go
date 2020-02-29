package seviper

import (
	"encoding/binary"
	"math"
)

// Writer for Write to Bytes
type Writer struct {
	Bytes   Bytes
	c       int
	isFixed bool
}

// NewWriter is Construct of Bytes
func NewWriter(opt ...interface{}) *Writer {
	if len(opt) > 0 {
		switch opt[0].(type) {
		case []byte:
			return &Writer{Bytes: opt[0].([]byte), c: 0, isFixed: true}
		}
	}
	return &Writer{Bytes: []byte{}, c: 0, isFixed: false}
}

// Write is Method for Write to Bytes
func (bw *Writer) Write(v interface{}) bool {
	var buff []byte
	switch v.(type) {
	case int:
		buff = make([]byte, 4)
		binary.LittleEndian.PutUint32(buff, uint32(v.(int)))
	case string:
		buff = []byte(v.(string) + "\000")
	case float32:
		buff = make([]byte, 4)
		binary.LittleEndian.PutUint32(buff[:], math.Float32bits(v.(float32)))
	case float64:
		buff = make([]byte, 8)
		binary.LittleEndian.PutUint64(buff[:], math.Float64bits(v.(float64)))
	}
	if !bw.isFixed {
		writesize := len(buff)
		if l := len(bw.Bytes) - bw.c; l < writesize {
			bw.Bytes = append(bw.Bytes, make([]byte, writesize-l)...)
		}
	}
	ok := NewBytes(bw.Bytes).Write(bw.c, buff)
	if ok {
		bw.c += len(buff)
	}
	return ok
}
