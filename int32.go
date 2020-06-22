package convert

import (
	"strconv"
)

// Int32 forces a conversion from an arbitrary value into an int32.
// If the value cannot be converted, then the default value for the type is used.
func Int32(value interface{}) int32 {

	result, _ := Int32Ok(value)
	return result
}

// Int32Ok tries to convert an arbitrary value into an integer
func Int32Ok(value interface{}) (int32, bool) {

	if value == nil {
		return int32(0), false
	}

	switch v := value.(type) {

	case int:
		return int32(v), true

	case int8:
		return int32(v), true

	case int16:
		return int32(v), true

	case int32:
		return int32(v), true

	case int64:
		return int32(v), true

	case float32:
		return int32(v), true

	case float64:
		return int32(v), true

	case *int:
		return int32(*v), true

	case *int8:
		return int32(*v), true

	case *int16:
		return int32(*v), true

	case *int32:
		return int32(*v), true

	case *int64:
		return int32(*v), true

	case *float32:
		return int32(*v), true

	case *float64:
		return int32(*v), true

	case string:
		result, err := strconv.ParseInt(v, 10, 32)

		if err != nil {
			return 0, false
		}

		return int32(result), true
	}

	return 0, false
}
