package convert

import (
	"strconv"
)

// Float64 forces a conversion from an arbitrary value into a float64.
// If the value cannot be converted, then the default value for the type is used.
func Float64(value interface{}) float64 {

	result, _ := Float64Ok(value)
	return result
}

// Float64Ok tries to convert an arbitrary value into an float64
func Float64Ok(value interface{}) (float64, bool) {

	if value == nil {
		return float64(0), false
	}

	switch v := value.(type) {

	case int:
		return float64(v), true

	case int8:
		return float64(v), true

	case int16:
		return float64(v), true

	case int32:
		return float64(v), true

	case int64:
		return float64(v), true

	case float32:
		return float64(v), true

	case float64:
		return float64(v), true

	case *int:
		return float64(*v), true

	case *int8:
		return float64(*v), true

	case *int16:
		return float64(*v), true

	case *int32:
		return float64(*v), true

	case *int64:
		return float64(*v), true

	case *float32:
		return float64(*v), true

	case *float64:
		return float64(*v), true

	case string:
		result, err := strconv.ParseFloat(v, 64)

		if err != nil {
			return 0, false
		}

		return result, true
	}

	return 0, false
}
