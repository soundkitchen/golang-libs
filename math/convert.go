package math

import (
	"fmt"
)

// convert numeric to float64.
func convertToFloat64(v interface{}) float64 {
	switch v.(type) {
	case int:
		return float64(v.(int))
	case int8:
		return float64(v.(int8))
	case int16:
		return float64(v.(int16))
	case int32:
		return float64(v.(int32))
	case int64:
		return float64(v.(int64))
	case uint:
		return float64(v.(uint))
	case uint8:
		return float64(v.(uint8))
	case uint16:
		return float64(v.(uint16))
	case uint32:
		return float64(v.(uint32))
	case uint64:
		return float64(v.(uint64))
	case float32:
		return float64(v.(float32))
	case float64:
		return v.(float64)
	default:
		panic(fmt.Sprintf("can't convert to float64 %s", v))
	}
}
