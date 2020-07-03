package convert

import (
	"strconv"
)

// Int32 forces a conversion from an arbitrary value into an int32.
// If the value cannot be converted, then the zero value for the type (0) is used.
func Int32(value interface{}, defaultValue int32) int32 {

	result, _ := Int32Natural(value, defaultValue)
	return result
}

// Int32Default forces a conversion from an arbitrary value into a int32.
// if the value cannot be converted, then the default value is used.
func Int32Default(value interface{}, defaultValue int32) int32 {

	result, _ := Int32Natural(value, defaultValue)
	return result
}

// Int32Natural converts an arbitrary value (passed in the first parameter) into an int32, no matter what.
// The first result is the final converted value, or the default value (passed in the second parameter)
// The second result is TRUE if the value was naturally an integer, and FALSE otherwise
func Int32Natural(value interface{}, defaultValue int32) (int32, bool) {

	if value == nil {
		return defaultValue, false
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
		return int32(v), false

	case float64:
		return int32(v), false

	case string:
		result, err := strconv.ParseInt(v, 10, 32)

		if err != nil {
			return defaultValue, false
		}

		return int32(result), false

	case Stringer:
		return Int32Natural(v.String(), defaultValue)
	}

	return defaultValue, false
}
