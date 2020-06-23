package convert

import (
	"strconv"
)

// Int forces a conversion from an arbitrary value into an int.
// If the value cannot be converted, then the default value for the type is used.
func Int(value interface{}, defaultValue int) int {

	result, _ := NaturalInt(value, defaultValue)
	return result
}

// NaturalInt converts an arbitrary value (passed in the first parameter) into an int, no matter what.
// The first result is the final converted value, or the default value (passed in the second parameter)
// The second result is TRUE if the value was naturally an integer, and FALSE otherwise
func NaturalInt(value interface{}, defaultValue int) (int, bool) {

	if value == nil {
		return defaultValue, false
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
		return int(v), false

	case float64:
		return int(v), false

	case string:
		result, err := strconv.Atoi(v)

		if err != nil {
			return defaultValue, false
		}

		return result, false

	case Stringer:
		return NaturalInt(v.String(), defaultValue)
	}

	return defaultValue, false
}
