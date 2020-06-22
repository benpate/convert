package convert

import (
	"strconv"
)

// Int forces a conversion from an arbitrary value into an int.
// If the value cannot be converted, then the default value for the type is used.
func Int(value interface{}, defaultValue int) int {

	if result, ok := IntOk(value); ok {
		return result
	}

	return defaultValue
}

// IntOk tries to convert an arbitrary value into an integer
func IntOk(value interface{}) (int, bool) {

	if value == nil {
		return int(0), false
	}

	switch v := value.(type) {

	case int:
		return int(v), true

	case int8:
		return int(v), true

	case int16:
		return int(v), true

	case int32:
		return int(v), true

	case int64:
		return int(v), true

	case float32:
		return int(v), true

	case float64:
		return int(v), true

	case *int:
		return int(*v), true

	case *int8:
		return int(*v), true

	case *int16:
		return int(*v), true

	case *int32:
		return int(*v), true

	case *int64:
		return int(*v), true

	case *float32:
		return int(*v), true

	case *float64:
		return int(*v), true

	case string:
		result, err := strconv.Atoi(v)

		if err != nil {
			return 0, false
		}

		return result, true
	}

	return 0, false
}
