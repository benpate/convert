package convert

import (
	"strconv"

	"github.com/benpate/derp"
)

// Int64 tries to convert an arbitrary value into an integer
func Int64(value interface{}) (int64, *derp.Error) {

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
			return 0, derp.New(500, "convert.Int64", "Not a valid int64", value, err)
		}

		return int64(result), nil
	}

	return 0, derp.New(500, "convert.Int64", "Cannot Convert to int64", value)
}
