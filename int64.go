package convert

import (
	"strconv"

	"github.com/benpate/derp"
)

// MustInt64 forces a conversion from an arbitrary value into an int64.
// If the value cannot be converted, then the default value for the type is used.
func MustInt64(value interface{}) int64 {

	result, _ := Int64(value)
	return result
}

// Int64 tries to convert an arbitrary value into an integer
func Int64(value interface{}) (int64, *derp.Error) {

	if value == nil {
		return int64(0), derp.New(500, "convert.Int64", "null pointer")
	}

	switch v := value.(type) {

	case int:
		return int64(v), nil

	case int8:
		return int64(v), nil

	case int16:
		return int64(v), nil

	case int32:
		return int64(v), nil

	case int64:
		return int64(v), nil

	case float32:
		return int64(v), nil

	case float64:
		return int64(v), nil

	case *int:
		return int64(*v), nil

	case *int8:
		return int64(*v), nil

	case *int16:
		return int64(*v), nil

	case *int32:
		return int64(*v), nil

	case *int64:
		return int64(*v), nil

	case *float32:
		return int64(*v), nil

	case *float64:
		return int64(*v), nil

	case string:
		result, err := strconv.ParseInt(v, 10, 64)

		if err != nil {
			return 0, derp.Wrap(err, "convert.Int64", "Not a valid int64", value)
		}

		return int64(result), nil
	}

	return 0, derp.New(500, "convert.Int64", "Cannot convert to int64", value)
}
