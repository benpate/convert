package convert

import (
	"strconv"
)

// Int64 forces a conversion from an arbitrary value into an int64.
// If the value cannot be converted, then the zero value for the type is used.
func Int64(value interface{}) int64 {

	result, _ := Int64Natural(value, 0)
	return result

}

// Int64Default forces a conversion from an arbitrary value into a int64.
// if the value cannot be converted, then the default value is used.
func Int64Default(value interface{}, defaultValue int64) int64 {

	result, _ := Int64Natural(value, defaultValue)
	return result
}

// Int64Natural converts an arbitrary value (passed in the first parameter) into an int64, no matter what.
// The first result is the final converted value, or the default value (passed in the second parameter)
// The second result is TRUE if the value was naturally an integer, and FALSE otherwise
func Int64Natural(value interface{}, defaultValue int64) (int64, bool) {

	if value == nil {
		return defaultValue, false
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
		return int64(v), false

	case float64:
		return int64(v), false

	case string:
		result, err := strconv.ParseInt(v, 10, 64)

		if err != nil {
			return defaultValue, false
		}

		return int64(result), false

	case Stringer:
		return Int64Natural(v.String(), defaultValue)
	}

	return defaultValue, false
}
