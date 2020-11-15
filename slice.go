package convert

func SliceOfString(value []interface{}) []string {

	result := make([]string, len(value))

	for index, v := range value {
		result[index] = String(v)
	}

	return result
}
