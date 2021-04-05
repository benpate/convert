package convert

func SliceOfString(value interface{}) []string {

	switch value := value.(type) {

	case []interface{}:
		result := make([]string, len(value))
		for index, v := range value {
			result[index] = String(v)
		}
		return result

	case []string:
		return value

	case string:
		return []string{value}
	}

	return make([]string, 0)
}
