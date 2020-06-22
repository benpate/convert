package convert

import (
	"strconv"
)

// Float32 forces a conversion from an arbitrary value into a float32.
// If the value cannot be converted, then the default value for the type is used.
func Float32(value interface{}, defaultValue float32) float32 {

	if result, ok := Float32Ok(value); ok {
		return result
	}

	return defaultValue
}

// Float32Ok tries to convert an arbitrary value into an float32
func Float32Ok(value interface{}) (float32, bool) {

	if value == nil {
		return float32(0), false
	}

	switch v := value.(type) {

	case int:
		return float32(v), true

	case int8:
		return float32(v), true

	case int16:
		return float32(v), true

	case int32:
		return float32(v), true

	case int64:
		return float32(v), true

	case float32:
		return float32(v), true

	case float64:
		return float32(v), true

	case *int:
		return float32(*v), true

	case *int8:
		return float32(*v), true

	case *int16:
		return float32(*v), true

	case *int32:
		return float32(*v), true

	case *int64:
		return float32(*v), true

	case *float32:
		return float32(*v), true

	case *float64:
		return float32(*v), true

	case string:
		result, err := strconv.ParseFloat(v, 32)

		if err != nil {
			return 0, false
		}

		return float32(result), true

	case Stringer:
		return Float32Ok(v.String())
	}

	return 0, false
}
