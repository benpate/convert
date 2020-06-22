package convert

import (
	"strconv"

	"github.com/benpate/derp"
)

// MustInt32 forces a conversion from an arbitrary value into an int32.
// If the value cannot be converted, then the default value for the type is used.
func MustInt32(value interface{}) int32 {

	result, _ := Int32(value)
	return result
}

// Int32 tries to convert an arbitrary value into an integer
func Int32(value interface{}) (int32, *derp.Error) {

	if value == nil {
		return int32(0), derp.New(500, "convert.Int32", "null pointer")
	}

	switch v := value.(type) {

	case int:
		return int32(v), nil

	case int8:
		return int32(v), nil

	case int16:
		return int32(v), nil

	case int32:
		return int32(v), nil

	case int64:
		return int32(v), nil

	case float32:
		return int32(v), nil

	case float64:
		return int32(v), nil

	case *int:
		return int32(*v), nil

	case *int8:
		return int32(*v), nil

	case *int16:
		return int32(*v), nil

	case *int32:
		return int32(*v), nil

	case *int64:
		return int32(*v), nil

	case *float32:
		return int32(*v), nil

	case *float64:
		return int32(*v), nil

	case string:
		result, err := strconv.ParseInt(v, 10, 32)

		if err != nil {
			return 0, derp.Wrap(err, "convert.Int32", "Not a valid integer", value)
		}

		return int32(result), nil
	}

	return 0, derp.New(500, "convert.Int32", "Cannot convert to integer", value)
}
