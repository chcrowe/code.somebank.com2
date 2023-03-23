package bytes

import (
	"bytes"
	"fmt"
)

type SafeBuffer struct {
	buf         []byte
	off         int
	granularity uint
}

func NewSafeBuffer(l uint) *SafeBuffer {
	return &SafeBuffer{buf: make([]byte, l), off: 0, granularity: l}
}

func NewSafeBufferFromExisting(b []byte, l uint) *SafeBuffer {
	return &SafeBuffer{buf: b, off: len(b) - 1, granularity: l}
}

func RoundUp(size uint, granularity uint) uint {

	if size < granularity {
		return granularity
	}

	times := size / granularity
	remainder := granularity % size

	newsize := times * granularity

	if 0 == remainder {
		return newsize
	} else {
		return newsize + granularity
	}
}

func (b *SafeBuffer) Zeroize() {
	for i, _ := range b.buf {
		b.buf[i] = 0
	}
	b.off = 0
}

func (b *SafeBuffer) Bytes() []byte { return b.buf[0:b.off] }

func (b *SafeBuffer) checkCapacity(dlen int) {

	newlen := dlen + b.off
	if newlen > cap(b.buf) {
		newcap := RoundUp(uint(newlen), b.granularity)
		//grow the buffer (increase capacity)
		newSlice := make([]byte, uint(newcap))
		copy(newSlice, b.buf)
		b.buf = newSlice
	}

}

func (b *SafeBuffer) Length() int {
	return b.off
}

func (b *SafeBuffer) Capacity() int {
	return len(b.buf)
}

func (b *SafeBuffer) AppendByte(c byte) {
	b.checkCapacity(1)
	b.buf[b.off] = c
	b.off++
}

//func copy(dst, src []Type) int

func (b *SafeBuffer) Write(c []byte) (n int, err error) {
	b.checkCapacity(len(c))
	m := copy(b.buf[b.off:], c)
	b.off += m
	return m, nil
}

func (b *SafeBuffer) AppendBytes(c []byte) {
	b.checkCapacity(len(c))
	m := copy(b.buf[b.off:], c)
	b.off += m
}

func (b *SafeBuffer) AppendString(s string) {
	b.checkCapacity(len(s))
	copy(b.buf[b.off:], s)
	b.off += len(s)
}

func (b *SafeBuffer) AppendFormat(format string, a ...interface{}) {
	formattedstr := fmt.Sprintf(format, a...)
	b.checkCapacity(len(formattedstr))
	m := copy(b.buf[b.off:], formattedstr)
	b.off += m
}

func (b *SafeBuffer) String() string {
	return string(b.buf[0:b.off])
}

func (b *SafeBuffer) EncodeParity(start int) []byte {

	stxstart := bytes.Index(b.buf[0:b.off], []byte{byte(StartOfText)})
	if -1 == stxstart {
		panic("EncodeBuffer(): Start sentinel framing (STX=2) byte not found")
	}
	stxstart += 1
	var lrc byte = 0x00
	lrc = CalculateLRC(b.buf, stxstart, b.off-stxstart)
	b.AppendByte(lrc)

	// set parity on each byte
	for i := start; i < b.off; i++ {
		b.buf[i] = SetBitParity(b.buf[i], true)
	}

	return b.buf[start:b.off]
}

func (b *SafeBuffer) DecodeParity(start int) []byte {

	// set parity on each byte
	for i := 0; i < b.off; i++ {
		if 128 < uint(b.buf[i]) {
			b.buf[i] = byte(uint(b.buf[i]) - 128)
		}
	}

	return b.buf[start:b.off]
}
