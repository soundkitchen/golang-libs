package math

import (
	"testing"
)

// case int -> float64
func TestConvertToFloat64FromInt(t *testing.T) {
	var a int = 10
	if !mustFloat64(convertToFloat64(a)) {
		t.Fatalf("convert failure.")
	}
}

// case int8 -> float64
func TestConvertToFloat64FromInt8(t *testing.T) {
	var a int8 = 10
	if !mustFloat64(convertToFloat64(a)) {
		t.Fatalf("convert failure.")
	}
}

// case int16 -> float64
func TestConvertToFloat64FromInt16(t *testing.T) {
	var a int16 = 10
	if !mustFloat64(convertToFloat64(a)) {
		t.Fatalf("convert failure.")
	}
}

// case int32 -> float64
func TestConvertToFloat64FromInt32(t *testing.T) {
	var a int32 = 10
	if !mustFloat64(convertToFloat64(a)) {
		t.Fatalf("convert failure.")
	}
}

// case int64 -> float64
func TestConvertToFloat64FromInt64(t *testing.T) {
	var a int64 = 10
	if !mustFloat64(convertToFloat64(a)) {
		t.Fatalf("convert failure.")
	}
}

// case uint -> float64
func TestConvertToFloat64FromUint(t *testing.T) {
	var a uint = 10
	if !mustFloat64(convertToFloat64(a)) {
		t.Fatalf("convert failure.")
	}
}

// case uint8 -> float64
func TestConvertToFloat64FromUint8(t *testing.T) {
	var a uint8 = 10
	if !mustFloat64(convertToFloat64(a)) {
		t.Fatalf("convert failure.")
	}
}

// case uint16 -> float64
func TestConvertToFloat64FromUint16(t *testing.T) {
	var a uint16 = 10
	if !mustFloat64(convertToFloat64(a)) {
		t.Fatalf("convert failure.")
	}
}

// case uint32 -> float64
func TestConvertToFloat64FromUint32(t *testing.T) {
	var a uint32 = 10
	if !mustFloat64(convertToFloat64(a)) {
		t.Fatalf("convert failure.")
	}
}

// case uint64 -> float64
func TestConvertToFloat64FromUint64(t *testing.T) {
	var a uint64 = 10
	if !mustFloat64(convertToFloat64(a)) {
		t.Fatalf("convert failure.")
	}
}

// case float32 -> float64
func TestConvertToFloat64FromFloat32(t *testing.T) {
	var a float32 = 10
	if !mustFloat64(convertToFloat64(a)) {
		t.Fatalf("convert failure.")
	}
}

// check value type.
// return true if type is float64 otherwise return false.
func mustFloat64(v interface{}) bool {
	switch v.(type) {
	case float64:
		return true
	default:
		return false
	}
}
