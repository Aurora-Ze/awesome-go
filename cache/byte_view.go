package cache

import "github.com/spf13/cast"

// ByteView only for read
type ByteView struct {
	data []byte
}

func (v ByteView) Len() uint64 {
	return cast.ToUint64(len(v.data))
}

func (v ByteView) String() string {
	return string(v.data)
}

func (v ByteView) ByteSlice() []byte {
	bytes := make([]byte, len(v.data))
	copy(bytes, v.data)
	return bytes
}
