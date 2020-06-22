package convert

import (
	"io"
	"strconv"
)

// String forces a conversion from an arbitrary value into an string.
// If the value cannot be converted, then the default value for the type is used.
func String(value interface{}, defaultValue string) string {

	if result, ok := StringOk(value); ok {
		return result
	}

	return defaultValue
}

// StringOk tries to convert an arbitrary value into a string.
// If the value provided cannot be converted into a string, this
// function returns the default string value, and FALSE.
func StringOk(value interface{}) (string, bool) {

	if value == nil {
		return "", false
	}

	switch v := value.(type) {

	case string:
		return v, true

	case *string:
		return *v, true

	case []byte:
		return string(v), true

	case io.Reader:
		var buffer []byte

		if _, err := v.Read(buffer); err != nil {
			return "", false
		}

		return string(buffer), true

	case int:
		return strconv.Itoa(v), true

	case int8:
		return strconv.FormatInt(int64(v), 10), true

	case int16:
		return strconv.FormatInt(int64(v), 10), true

	case int32:
		return strconv.FormatInt(int64(v), 10), true

	case int64:
		return strconv.FormatInt(v, 10), true

	default:
		return "", false
	}
}
