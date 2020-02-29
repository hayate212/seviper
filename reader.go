package seviper

// Reader for Read to Bytes
type Reader struct {
	Bytes Bytes
	c     int
}

// NewReader is Construct of Reader
func NewReader(buff []byte) *Reader {
	br := Reader{Bytes: buff, c: 0}
	return &br
}

// ToInt reads and return int from Bytes
func (br *Reader) ToInt() int {
	value := NewBytes(br.Bytes[br.c : br.c+4]).ToInt()
	br.c += 4
	return value
}

// ToString reads and return string from Bytes
func (br *Reader) ToString() string {
	buff := br.Bytes[br.c:]
	str := buff.NULTrim().ToString()
	br.c += len(str) + 1
	return str
}

// ToFloat32 reads and return float32 from Bytes
func (br *Reader) ToFloat32() float32 {
	value := NewBytes(br.Bytes[br.c : br.c+4]).ToFloat32()
	br.c += 4
	return value
}

// ToFloat64 reads and return float64 from Bytes
func (br *Reader) ToFloat64() float64 {
	value := NewBytes(br.Bytes[br.c : br.c+8]).ToFloat64()
	br.c += 8
	return value
}

func (br *Reader) Seek(n int) {
	br.c += n
}
