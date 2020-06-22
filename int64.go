package convert

import (
	"strconv"
)

// Int64 forces a conversion from an arbitrary value into an int64.
// If the value cannot be converted, then the default value for the type is used.
func Int64(value interface{}, defaultValue int64) int64 {

	if result, ok := Int64Ok(value); ok {
		return result
	}

	return defaultValue
}

// Int64Ok tries to convert an arbitrary value into an integer
func Int64Ok(value interface{}) (int64, bool) {

	if value == nil {
		return int64(0), false
	}

	switch v := value.(type) {

	case int:
		return int64(v), true

	case int8:
		return int64(v), true

	case int16:
		return int64(v), true

	case int32:
		return int64(v), true

	case int64:
		return int64(v), true

	case float32:
		return int64(v), true

	case float64:
		return int64(v), true

	case *int:
		return int64(*v), true

	case *int8:
		return int64(*v), true

	case *int16:
		return int64(*v), true

	case *int32:
		return int64(*v), true

	case *int64:
		return int64(*v), true

	case *float32:
		return int64(*v), true

	case *float64:
		return int64(*v), true

	case string:
		result, err := strconv.ParseInt(v, 10, 64)

		if err != nil {
			return 0, false
		}

		return int64(result), true
	}

	return 0, false
}
