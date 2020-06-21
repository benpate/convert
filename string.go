package convert

import (
	"strconv"

	"github.com/benpate/derp"
)

// String tries to convert an arbitrary value into a string
func String(value interface{}) (string, *derp.Error) {

	switch v := value.(type) {
	case string:
		return v, nil

	case *string:
		return *v, nil

	case []byte:
		return string(v), nil

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
