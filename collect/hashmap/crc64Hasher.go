package hashmap

import (
	"fmt"
	"hash/crc64"
)

var gCrc64Table = crc64.MakeTable(crc64.ECMA)

type FnEquals func(a interface{}, b interface{}) bool

type tCrc64Hasher struct {
	fnEquals FnEquals
}

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX
const INT32_MAX = int32(^uint32(0) >> 1)
const INT32_MIN = ^INT32_MAX
const INT64_MAX = int64(^uint64(0) >> 1)
const INT64_MIN = ^INT64_MAX

func (me *tCrc64Hasher) Hash(it interface{}) uint64 {
	if it == nil {
		return 0
	}

	if v, ok := it.(int); ok {
		return uint64(v - INT_MIN)

	} else if v, ok := it.(int64); ok {
		return uint64(v - INT64_MIN)

	} else if v, ok := it.(int32); ok {
		return uint64(v - INT32_MIN)

	} else if v, ok := it.(uint32); ok {
		return uint64(v)

	} else if v, ok := it.(uint64); ok {
		return v

	} else if v, ok := it.(string); ok {
		return crc64.Checksum([]byte(v), gCrc64Table)

	} else {
		data := []byte(fmt.Sprintf("%v", it))
		return crc64.Checksum(data, gCrc64Table)
	}
}

func (me *tCrc64Hasher) Equals(a interface{}, b interface{}) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	fn := me.fnEquals
	if fn == nil {
		return a == b
	} else {
		return fn(a, b)
	}
}

func NewCrc64Hashful(fn FnEquals) IHasher {
	return &tCrc64Hasher{
		fnEquals: fn,
	}
}
