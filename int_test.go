package safeconversion_test

import (
	"errors"
	"math"
	"testing"

	"github.com/corentings/safeconversion"
)

func TestCastInt(t *testing.T) {
	t.Run("int to uint", testCastInt[int, uint](42, 42, false, nil))
	t.Run("uint to int", testCastInt[uint, int](42, 42, false, nil))
	t.Run("int8 to int16", testCastInt[int8, int16](127, 127, false, nil))
	t.Run("uint16 to uint32", testCastInt[uint16, uint32](65535, 65535, false, nil))
	t.Run("int32 to int64", testCastInt[int32, int64](2147483647, 2147483647, false, nil))
	t.Run("uint64 to uint", testCastInt[uint64, uint](18446744073709551615, 18446744073709551615, false, nil))
	t.Run("int to int8 (error)", testCastInt[int, int8](300, 0, true, safeconversion.ErrValueOutOfRange))
	t.Run("uint to int8 (error)", testCastInt[uint, int8](300, 0, true, safeconversion.ErrValueOutOfRange))
	t.Run("int8 to uint16 (no error)", testCastInt[int8, uint16](127, 127, false, nil))
	t.Run("uint16 to int8 (overflow)", testCastInt[uint16, int8](256, 0, true, safeconversion.ErrValueOutOfRange))
	t.Run("uint16 to int8 (in range)", testCastInt[uint16, int8](127, 127, false, nil))
	t.Run("int8 to uint16 (error)", testCastInt[int8, uint16](-128, 0, true, safeconversion.ErrValueOutOfRange))

	// Additional boundary cases
	t.Run("int8 max to uint8", testCastInt[int8, uint8](127, 127, false, nil))
	t.Run("int8 min to uint8", testCastInt[int8, uint8](-128, 0, true, safeconversion.ErrValueOutOfRange))
	t.Run("uint8 max to int8", testCastInt[uint8, int8](255, 0, true, safeconversion.ErrValueOutOfRange))
	t.Run("int16 max to uint16", testCastInt[int16, uint16](32767, 32767, false, nil))
	t.Run("int16 min to uint16", testCastInt[int16, uint16](-32768, 0, true, safeconversion.ErrValueOutOfRange))
	t.Run("uint16 max to int16", testCastInt[uint16, int16](65535, 0, true, safeconversion.ErrValueOutOfRange))
	t.Run("int32 max to uint32", testCastInt[int32, uint32](2147483647, 2147483647, false, nil))
	t.Run("int32 min to uint32", testCastInt[int32, uint32](-2147483648, 0, true, safeconversion.ErrValueOutOfRange))
	t.Run("uint32 max to int32", testCastInt[uint32, int32](4294967295, 0, true, safeconversion.ErrValueOutOfRange))
	t.Run("int64 max to uint64", testCastInt[int64, uint64](9223372036854775807, 9223372036854775807, false, nil))
	t.Run("int64 min to uint64", testCastInt[int64, uint64](-9223372036854775808, 0, true, safeconversion.ErrValueOutOfRange))
	t.Run("uint64 max to int64", testCastInt[uint64, int64](18446744073709551615, 0, true, safeconversion.ErrValueOutOfRange))
	t.Run("int to uint8 (max)", testCastInt[int, uint8](math.MaxUint8, math.MaxUint8, false, nil))
	t.Run("int to uint8 (overflow)", testCastInt[int, uint8](math.MaxUint8+1, 0, true, safeconversion.ErrValueOutOfRange))
	t.Run("int to int8 (min)", testCastInt[int, int8](math.MinInt8, math.MinInt8, false, nil))
	t.Run("int to int8 (underflow)", testCastInt[int, int8](-129, 0, true, safeconversion.ErrValueOutOfRange))
	t.Run("uint to int (max 32-bit)", testCastInt[uint, int](math.MaxInt32, math.MaxInt32, false, nil))
	t.Run("int64 to uint32 (max)", testCastInt[int64, uint32](math.MaxUint32, math.MaxUint32, false, nil))
	t.Run("int64 to uint32 (overflow)", testCastInt[int64, uint32](math.MaxUint32+1, 0, true, safeconversion.ErrValueOutOfRange))
}

func FuzzCastInt(f *testing.F) {
	f.Add(int(0))
	f.Add(int(1))
	f.Add(int(-1))
	f.Add(int(math.MaxInt64))
	f.Add(int(math.MinInt64))

	f.Fuzz(func(t *testing.T, input int) {
		testCastIntFuzz[int, int](t, input)
		testCastIntFuzz[int, int8](t, input)
		testCastIntFuzz[int, int16](t, input)
		testCastIntFuzz[int, int32](t, input)
		testCastIntFuzz[int, int64](t, input)
		testCastUintFuzz[int, uint](t, input)
		testCastUintFuzz[int, uint8](t, input)
		testCastUintFuzz[int, uint16](t, input)
		testCastUintFuzz[int, uint32](t, input)
		testCastUintFuzz[int, uint64](t, input)
	})
}

func FuzzCastInt32(f *testing.F) {
	f.Add(int32(0))
	f.Add(int32(1))
	f.Add(int32(-1))
	f.Add(int32(math.MaxInt32))
	f.Add(int32(math.MinInt32))

	f.Fuzz(func(t *testing.T, input int32) {
		testCastIntFuzz[int32, int](t, input)
		testCastIntFuzz[int32, int8](t, input)
		testCastIntFuzz[int32, int16](t, input)
		testCastIntFuzz[int32, int32](t, input)
		testCastIntFuzz[int32, int64](t, input)
		testCastUintFuzz[int32, uint](t, input)
		testCastUintFuzz[int32, uint8](t, input)
		testCastUintFuzz[int32, uint16](t, input)
		testCastUintFuzz[int32, uint32](t, input)
		testCastUintFuzz[int32, uint64](t, input)
	})
}

func FuzzCastInt64(f *testing.F) {
	f.Add(int64(0))
	f.Add(int64(1))
	f.Add(int64(-1))
	f.Add(int64(math.MaxInt64))
	f.Add(int64(math.MinInt64))

	f.Fuzz(func(t *testing.T, input int64) {
		testCastIntFuzz[int64, int](t, input)
		testCastIntFuzz[int64, int8](t, input)
		testCastIntFuzz[int64, int16](t, input)
		testCastIntFuzz[int64, int32](t, input)
		testCastIntFuzz[int64, int64](t, input)
		testCastUintFuzz[int64, uint](t, input)
		testCastUintFuzz[int64, uint8](t, input)
		testCastUintFuzz[int64, uint16](t, input)
		testCastUintFuzz[int64, uint32](t, input)
		testCastUintFuzz[int64, uint64](t, input)
	})
}

func FuzzCastUint(f *testing.F) {
	f.Add(uint(0))
	f.Add(uint(1))
	f.Add(uint(math.MaxUint64))

	f.Fuzz(func(t *testing.T, input uint) {
		testCastIntFuzz[uint, uint](t, input)
		testCastIntFuzz[uint, uint8](t, input)
		testCastIntFuzz[uint, uint16](t, input)
		testCastIntFuzz[uint, uint32](t, input)
		testCastIntFuzz[uint, uint64](t, input)
		testCastIntFuzz[uint, int](t, input)
		testCastIntFuzz[uint, int8](t, input)
		testCastIntFuzz[uint, int16](t, input)
		testCastIntFuzz[uint, int32](t, input)
		testCastIntFuzz[uint, int64](t, input)
	})
}

func FuzzCastUint8(f *testing.F) {
	f.Add(uint8(0))
	f.Add(uint8(1))
	f.Add(uint8(math.MaxUint8))

	f.Fuzz(func(t *testing.T, input uint8) {
		testCastIntFuzz[uint8, uint](t, input)
		testCastIntFuzz[uint8, uint8](t, input)
		testCastIntFuzz[uint8, uint16](t, input)
		testCastIntFuzz[uint8, uint32](t, input)
		testCastIntFuzz[uint8, uint64](t, input)
		testCastIntFuzz[uint8, int](t, input)
		testCastIntFuzz[uint8, int8](t, input)
		testCastIntFuzz[uint8, int16](t, input)
		testCastIntFuzz[uint8, int32](t, input)
		testCastIntFuzz[uint8, int64](t, input)
	})
}

func FuzzCastUint16(f *testing.F) {
	f.Add(uint16(0))
	f.Add(uint16(1))
	f.Add(uint16(math.MaxUint16))

	f.Fuzz(func(t *testing.T, input uint16) {
		testCastIntFuzz[uint16, uint](t, input)
		testCastIntFuzz[uint16, uint8](t, input)
		testCastIntFuzz[uint16, uint16](t, input)
		testCastIntFuzz[uint16, uint32](t, input)
		testCastIntFuzz[uint16, uint64](t, input)
		testCastIntFuzz[uint16, int](t, input)
		testCastIntFuzz[uint16, int8](t, input)
		testCastIntFuzz[uint16, int16](t, input)
		testCastIntFuzz[uint16, int32](t, input)
		testCastIntFuzz[uint16, int64](t, input)
	})
}

func FuzzCastUint32(f *testing.F) {
	f.Add(uint32(0))
	f.Add(uint32(1))
	f.Add(uint32(math.MaxUint32))

	f.Fuzz(func(t *testing.T, input uint32) {
		testCastIntFuzz[uint32, uint](t, input)
		testCastIntFuzz[uint32, uint8](t, input)
		testCastIntFuzz[uint32, uint16](t, input)
		testCastIntFuzz[uint32, uint32](t, input)
		testCastIntFuzz[uint32, uint64](t, input)
		testCastIntFuzz[uint32, int](t, input)
		testCastIntFuzz[uint32, int8](t, input)
		testCastIntFuzz[uint32, int16](t, input)
		testCastIntFuzz[uint32, int32](t, input)
		testCastIntFuzz[uint32, int64](t, input)
	})
}

func FuzzCastUint64(f *testing.F) {
	f.Add(uint64(0))
	f.Add(uint64(1))
	f.Add(uint64(math.MaxUint64))

	f.Fuzz(func(t *testing.T, input uint64) {
		testCastIntFuzz[uint64, uint](t, input)
		testCastIntFuzz[uint64, uint8](t, input)
		testCastIntFuzz[uint64, uint16](t, input)
		testCastIntFuzz[uint64, uint32](t, input)
		testCastIntFuzz[uint64, uint64](t, input)
		testCastIntFuzz[uint64, int](t, input)
		testCastIntFuzz[uint64, int8](t, input)
		testCastIntFuzz[uint64, int16](t, input)
		testCastIntFuzz[uint64, int32](t, input)
		testCastIntFuzz[uint64, int64](t, input)
	})
}

func testCastUintFuzz[From, To safeconversion.Integer](t *testing.T, input From) {
	expectError := false
	if input < 0 {
		expectError = true
	}

	result, err := safeconversion.CastInt[From, To](input)

	if expectError && err == nil {
		t.Errorf("CastInt(%v) = %v, want error, got nil", input, result)

	} else if err != nil {
		if !errors.Is(err, safeconversion.ErrValueOutOfRange) {
			t.Errorf("Unexpected error type: %v", err)
		}
	}

	if err == nil && From(result) != input {
		t.Errorf("CastInt(%v) = %v, want %v", input, result, input)
	}
}

func testCastIntFuzz[From, To safeconversion.Integer](t *testing.T, input From) {
	result, err := safeconversion.CastInt[From, To](input)

	if err != nil {
		// If there's an error, make sure it's one of the expected error types
		if !errors.Is(err, safeconversion.ErrValueOutOfRange) {
			t.Errorf("Unexpected error type: %v", err)
		}
	} else {
		// If there's no error, verify that the result is within the valid range for the To type
		if From(result) != input {
			t.Errorf("Result %v does not match input %v", result, input)
		}
	}
}

func testCastInt[From, To safeconversion.Integer](input From, expected To, expectError bool, errorMsg error) func(*testing.T) {
	return func(t *testing.T) {
		result, err := safeconversion.CastInt[From, To](input)
		if expectError && err == nil {
			t.Errorf("CastInt(%v) = %v, want error, got nil", input, result)
		} else if !expectError && err != nil {
			t.Errorf("CastInt(%v) = error, want %v, got %v", input, errorMsg, err)
		} else if result != expected {
			t.Errorf("CastInt(%v) = %v, want %v", input, result, expected)
		} else if expectError && err != nil && !errors.Is(err, errorMsg) {
			t.Errorf("CastInt(%v) = error, want %v, got %v", input, errorMsg, err)
		}
	}
}
