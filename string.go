package convert

import "github.com/benpate/derp"

// String tries to convert an arbitrary value into a string
func String(value interface{}) (string, *derp.Error) {

	switch v := value.(type) {
	case string:
		return v, nil

	case *string:
		return *v, nil

	case []byte:
		return string(v), nil

	default:
		return "", derp.New(500, "convert.String", "Not a valid string", value)
	}
}
