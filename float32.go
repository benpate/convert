package convert

import (
	"strconv"

	"github.com/benpate/derp"
)

// Float32 tries to convert an arbitrary value into an float32
func Float32(value interface{}) (float32, *derp.Error) {

	switch v := value.(type) {

	case int:
		return float32(v), nil

	case int8:
		return float32(v), nil

	case int16:
		return float32(v), nil

	case int32:
		return float32(v), nil

	case int64:
		return float32(v), nil

	case float32:
		return float32(v), nil

	case float64:
		return float32(v), nil

	case *int:
		return float32(*v), nil

	case *int8:
		return float32(*v), nil

	case *int16:
		return float32(*v), nil

	case *int32:
		return float32(*v), nil

	case *int64:
		return float32(*v), nil

	case *float32:
		return float32(*v), nil

	case *float64:
		return float32(*v), nil

	case string:
		result, err := strconv.ParseFloat(v, 32)

		if err != nil {
			return 0, derp.Wrap(err, "convert.Float32", "Not a valid float32", value)
		}

		return float32(result), nil
	}

	return 0, derp.New(500, "convert.Float32", "Cannot Convert to float32", value)
}
