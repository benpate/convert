package convert

import (
	"strconv"
)

// Stringer enables a custom type to facilitate conversions to Strings.
type Stringer interface {
	String() string
}

// String forces a conversion from an arbitrary value into an string.
// If the value cannot be converted, then the default value for the type is used.
func String(value interface{}) string {

	result, _ := NaturalString(value, "")
	return result
}

// StringDefault forces a conversion from an arbitrary value into a string.
// if the value cannot be converted, then the default value is used.
func StringDefault(value interface{}, defaultValue string) string {

	result, _ := NaturalString(value, defaultValue)
	return result
}

// NaturalString converts an arbitrary value (passed in the first parameter) into a string, no matter what.
// The first result is the final converted value, or the default value (passed in the second parameter)
// The second result is TRUE if the value was naturally a string, and FALSE otherwise
func NaturalString(value interface{}, defaultValue string) (string, bool) {

	if value == nil {
		return "", false
	}

	switch v := value.(type) {

	case string:
		return v, true

	case Stringer:
		return v.String(), false

	case []byte:
		return string(v), true

	case int:
		return strconv.Itoa(v), false

	case int8:
		return strconv.FormatInt(int64(v), 10), false

	case int16:
		return strconv.FormatInt(int64(v), 10), false

	case int32:
		return strconv.FormatInt(int64(v), 10), false

	case int64:
		return strconv.FormatInt(v, 10), false

	case float32:
		return strconv.FormatFloat(float64(v), 'f', -2, 64), false

	case float64:
		return strconv.FormatFloat(v, 'f', -2, 64), false
	}

	return defaultValue, false
}
