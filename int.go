package convert

import (
	"strconv"

	"github.com/benpate/derp"
)

// MustInt forces a conversion from an arbitrary value into an int.
// If the value cannot be converted, then the default value for the type is used.
func MustInt(value interface{}) int {

	result, _ := Int(value)
	return result
}

// Int tries to convert an arbitrary value into an integer
func Int(value interface{}) (int, *derp.Error) {

	if value == nil {
		return int(0), derp.New(500, "convert.Int", "null pointer")
	}

	switch v := value.(type) {

	case int:
		return int(v), nil

	case int8:
		return int(v), nil

	case int16:
		return int(v), nil

	case int32:
		return int(v), nil

	case int64:
		return int(v), nil

	case float32:
		return int(v), nil

	case float64:
		return int(v), nil

	case *int:
		return int(*v), nil

	case *int8:
		return int(*v), nil

	case *int16:
		return int(*v), nil

	case *int32:
		return int(*v), nil

	case *int64:
		return int(*v), nil

	case *float32:
		return int(*v), nil

	case *float64:
		return int(*v), nil

	case string:
		result, err := strconv.Atoi(v)

		if err != nil {
			return 0, derp.Wrap(err, "convert.Int", "Not a valid integer", value)
		}

		return result, nil
	}

	return 0, derp.New(500, "convert.Int", "Cannot convert to Integer", value)
}
