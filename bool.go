package convert

import (
	"strconv"
)

// Bool forces a conversion from an arbitrary value into a boolean.
// If the value cannot be converted, then the default value for the type is used.
func Bool(value interface{}, defaultValue bool) bool {

	if result, ok := BoolOk(value); ok {
		return result
	}

	return defaultValue
}

// BoolOk tries to convert an arbitrary value into an bool
func BoolOk(value interface{}) (bool, bool) {

	if value == nil {
		return false, false
	}

	switch v := value.(type) {

	case bool:
		return v, true

	case int8:
		return (v != 0), true

	case int16:
		return (v != 0), true

	case int32:
		return (v != 0), true

	case int64:
		return (v != 0), true

	case float32:
		return (v != 0), true

	case float64:
		return (v != 0), true

	case *int:
		return (*v != 0), true

	case *int8:
		return (*v != 0), true

	case *int16:
		return (*v != 0), true

	case *int32:
		return (*v != 0), true

	case *int64:
		return (*v != 0), true

	case *float32:
		return (*v != 0), true

	case *float64:
		return (*v != 0), true

	case string:

		switch v {
		case "true", "yes":
			return true, true
		case "false", "no":
			return false, true
		}

		if result, err := strconv.ParseFloat(v, 64); err == nil {
			return BoolOk(result)
		}

	case Stringer:
		return BoolOk(v.String())
	}

	return false, false
}
