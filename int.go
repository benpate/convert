package convert

import (
	"strconv"
)

// Int forces a conversion from an arbitrary value into an int.
// If the value cannot be converted, then the zero value for the type (0) is used.
func Int(value interface{}) int {

	result, _ := IntNatural(value, 0)
	return result
}

// IntDefault forces a conversion from an arbitrary value into a int.
// if the value cannot be converted, then the default value is used.
func IntDefault(value interface{}, defaultValue int) int {

	result, _ := IntNatural(value, defaultValue)
	return result
}

// IntNatural converts an arbitrary value (passed in the first parameter) into an int, no matter what.
// The first result is the final converted value, or the default value (passed in the second parameter)
// The second result is TRUE if the value was naturally an integer, and FALSE otherwise
func IntNatural(value interface{}, defaultValue int) (int, bool) {

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
		return IntNatural(v.String(), defaultValue)
	}

	return defaultValue, false
}
