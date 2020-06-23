package convert

// Bool forces a conversion from an arbitrary value into a boolean.
// If the value cannot be converted, then the default value for the type is used.
func Bool(value interface{}, defaultValue bool) bool {

	result, _ := NaturalBool(value, defaultValue)
	return result
}

// NaturalBool converts an arbitrary value (passed in the first parameter) into a boolean, somehow, no matter what.
// The first result is the final converted value, or the default value (passed in the second parameter)
// The second result is TRUE if the value was naturally a bool, and FALSE otherwise
func NaturalBool(value interface{}, defaultValue bool) (bool, bool) {

	if value == nil {
		return defaultValue, false
	}

	switch v := value.(type) {

	case bool:
		return v, true

	case int:
		return (v != 0), false

	case int8:
		return (v != 0), false

	case int16:
		return (v != 0), false

	case int32:
		return (v != 0), false

	case int64:
		return (v != 0), false

	case float32:
		return (v != 0), false

	case float64:
		return (v != 0), false

	case string:

		switch v {
		case "true":
			return true, false
		case "false":
			return false, false
		}

	case Stringer:
		return NaturalBool(v.String(), defaultValue)
	}

	return defaultValue, false
}
