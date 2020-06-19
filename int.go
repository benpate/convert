package convert

import (
	"strconv"

	"github.com/benpate/derp"
)

// Int tries to convert an arbitrary value into an integer
func Int(value interface{}) (int, *derp.Error) {

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
			return 0, derp.New(500, "convert.Int", "Not a valid integer", value, err)
		}

		return result, nil
	}

	return 0, derp.New(500, "convert.Int", "Cannot Convert to Integer", value)
}
