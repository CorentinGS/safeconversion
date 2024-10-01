package safeconversion

type Integer interface {
	~int | ~uint | ~int8 | ~uint8 | ~int16 | ~uint16 | ~int32 | ~uint32 | ~int64 | ~uint64
}

func CastInt[From Integer, To Integer](value From) (To, error) {
	fromPositive := value >= 0

	converted := To(value)
	if fromPositive != (converted >= 0) {
		return 0, ErrValueOutOfRange
	}

	if From(converted) != value {
		return 0, ErrValueOutOfRange
	}

	return converted, nil
}
