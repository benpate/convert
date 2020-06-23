package convert

import (
	"strconv"
)

// Float64 forces a conversion from an arbitrary value into a float64.
// If the value cannot be converted, then the zero value for the type (false) is used.
func Float64(value interface{}) float64 {

	result, _ := NaturalFloat64(value, 0)
	return result
}

// Float64Default forces a conversion from an arbitrary value into a float64.
// if the value cannot be converted, then the default value is used.
func Float64Default(value interface{}, defaultValue float64) float64 {

	result, _ := NaturalFloat64(value, defaultValue)
	return result
}

// NaturalFloat64 converts an arbitrary value (passed in the first parameter) into a float64, no matter what.
// The first result is the final converted value, or the default value (passed in the second parameter)
// The second result is TRUE if the value was naturally a floating point number, and FALSE otherwise
func NaturalFloat64(value interface{}, defaultValue float64) (float64, bool) {

	if value == nil {
		return defaultValue, false
	}

	switch v := value.(type) {

	case int:
		return float64(v), false

	case int8:
		return float64(v), false

	case int16:
		return float64(v), false

	case int32:
		return float64(v), false

	case int64:
		return float64(v), false

	case float32:
		return float64(v), true

	case float64:
		return float64(v), true

	case string:
		result, err := strconv.ParseFloat(v, 64)

		if err != nil {
			return defaultValue, false
		}

		return result, false

	case Stringer:
		return NaturalFloat64(v.String(), defaultValue)
	}

	return defaultValue, false
}
