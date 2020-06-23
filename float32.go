package convert

import (
	"strconv"
)

// Float32 forces a conversion from an arbitrary value into a float32.
// If the value cannot be converted, then the default value for the type is used.
func Float32(value interface{}, defaultValue float32) float32 {

	result, _ := NaturalFloat32(value, defaultValue)
	return result
}

// NaturalFloat32 converts an arbitrary value (passed in the first parameter) into a float32, somehow, no matter what.
// The first result is the final converted value, or the default value (passed in the second parameter)
// The second result is TRUE if the value was naturally a floating point number, and FALSE otherwise
func NaturalFloat32(value interface{}, defaultValue float32) (float32, bool) {

	if value == nil {
		return defaultValue, false
	}

	switch v := value.(type) {

	case int:
		return float32(v), false

	case int8:
		return float32(v), false

	case int16:
		return float32(v), false

	case int32:
		return float32(v), false

	case int64:
		return float32(v), false

	case float32:
		return float32(v), true

	case float64:
		return float32(v), true

	case string:
		result, err := strconv.ParseFloat(v, 32)

		if err != nil {
			return defaultValue, false
		}

		return float32(result), false

	case Stringer:
		return NaturalFloat32(v.String(), defaultValue)
	}

	return defaultValue, false
}
