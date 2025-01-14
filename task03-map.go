package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {

	keys := make([]int, len(input))

	for k := range input {

		keys[k] = k

	}

	sort.Ints(keys)

	for _, k := range keys {

		result = append(result, input[k])
	}
	return result
}
