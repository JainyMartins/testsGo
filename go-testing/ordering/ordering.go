package ordering

import "sort"

func OrderArray(array []int) []int{
	sort.Ints(array)
	return array
}