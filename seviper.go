package seviper

import (
	"encoding/binary"
	"unsafe"
)

// Bytes []byte
type Bytes []byte

// NewBytes Bytes Construct
func NewBytes(buff []byte) *Bytes {
	var bytes Bytes
	bytes = buff
	return &bytes
}

// ToInt reads and return int from bytes
func (buff *Bytes) ToInt() int {
	return (int)(binary.LittleEndian.Uint32(*buff))
}

// ToString reads and return string from bytes
func (buff *Bytes) ToString() string {
	return *(*string)(unsafe.Pointer(buff))
}

// ToFloat32 reads and return float32 from bytes
func (buff *Bytes) ToFloat32() float32 {
	return *(*float32)(unsafe.Pointer(&(*buff)[0]))
}

// ToFloat64 reads and return float32 from bytes
func (buff *Bytes) ToFloat64() float64 {
	return *(*float64)(unsafe.Pointer(&(*buff)[0]))
}

// Write Bytes
func (buff *Bytes) Write(s int, t Bytes) bool {
	if s < 0 || len(*buff) < (s+len(t)) {
		return false
	}
	for i := 0; i < len(t); i++ {
		(*buff)[s+i] = t[i]
	}
	return true
}

// NULTrim is trim to Zero(0x0)
func (buff *Bytes) NULTrim() *Bytes {
	i := 0
	for ; (*buff)[i] != 0; i++ {
	}
	return NewBytes((*buff)[0:i])
}
