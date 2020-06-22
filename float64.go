package convert

import (
	"strconv"

	"github.com/benpate/derp"
)

// Float64 tries to convert an arbitrary value into an float64
func Float64(value interface{}) (float64, *derp.Error) {

	if value == nil {
		return float64(0), derp.New(500, "convert.Float64", "null pointer")
	}

	switch v := value.(type) {

	case int:
		return float64(v), nil

	case int8:
		return float64(v), nil

	case int16:
		return float64(v), nil

	case int32:
		return float64(v), nil

	case int64:
		return float64(v), nil

	case float32:
		return float64(v), nil

	case float64:
		return float64(v), nil

	case *int:
		return float64(*v), nil

	case *int8:
		return float64(*v), nil

	case *int16:
		return float64(*v), nil

	case *int32:
		return float64(*v), nil

	case *int64:
		return float64(*v), nil

	case *float32:
		return float64(*v), nil

	case *float64:
		return float64(*v), nil

	case string:
		result, err := strconv.ParseFloat(v, 64)

		if err != nil {
			return 0, derp.Wrap(err, "convert.Float64", "Not a valid float64", value)
		}

		return result, nil
	}

	return 0, derp.New(500, "convert.Float64", "Cannot convert to float64", value)
}
