package util

import (
	"math"
)

type F32 = float32
type F64 = float64

type I32 = int32
type I64 = int64

type U32 = uint32
type U64 = uint64

func Round(x F32) F32 {
	return F32(math.Round(F64(x)))
}

func BoolI32(x bool) I32 {
	var result I32
	if x {
		result = 1
	}
	return result
}
