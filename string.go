package safeconversion

import (
	"errors"
	"strconv"
)

var ErrOutOfRange = errors.New("value out of range for target type")

type numericType interface {
	~int | ~int32 | ~int64 | ~uint32 | ~uint64
}

func SafeParse[T numericType](value string) (T, error) { //nolint: ireturn // False positive
	var bitSize int
	switch any(T(0)).(type) {
	case int8, uint8:
		bitSize = 8
	case int16, uint16:
		bitSize = 16
	case int32, uint32:
		bitSize = 32
	case int64, uint64:
		bitSize = 64
	default:
		bitSize = 64
	}
	numericValue, err := strconv.ParseInt(value, 10, bitSize)
	if err != nil {
		return T(0), err
	}

	return T(numericValue), nil
}
