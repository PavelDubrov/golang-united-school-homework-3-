package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	var my_slice []int
	for k, _ := range input {
		my_slice = append(my_slice, k)
	}
	sort.Ints(my_slice)
	for _, elem := range my_slice {
		result = append(result, input[elem])
	}
	return result
}
