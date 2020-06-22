package convert

import (
	"io"
	"strconv"

	"github.com/benpate/derp"
)

// MustString forces a conversion from an arbitrary value into an string.
// If the value cannot be converted, then the default value for the type is used.
func MustString(value interface{}) string {

	result, _ := String(value)
	return result
}

// String tries to convert an arbitrary value into a string
func String(value interface{}) (string, *derp.Error) {

	if value == nil {
		return "", derp.New(500, "convert.String", "null pointer")
	}

	switch v := value.(type) {

	case string:
		return v, nil

	case *string:
		return *v, nil

	case []byte:
		return string(v), nil

	case io.Reader:
		var buffer []byte

		if _, err := v.Read(buffer); err != nil {
			return "", derp.Wrap(err, "convert.String", "Error reading io.Reader")
		}

		return string(buffer), nil

	case int:
		return strconv.Itoa(v), nil

	case int8:
		return strconv.FormatInt(int64(v), 10), nil

	case int16:
		return strconv.FormatInt(int64(v), 10), nil

	case int32:
		return strconv.FormatInt(int64(v), 10), nil

	case int64:
		return strconv.FormatInt(v, 10), nil

	default:
		return "", derp.New(500, "convert.String", "Not a valid string", value)
	}
}
