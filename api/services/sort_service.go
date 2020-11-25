package services

import "github.com/beltranbot/go-benchmarking/api/utils/sort"

// Sort func
func Sort(elements []int) {
	if len(elements) <= 10000 {
		sort.BubbleSort(elements)
		return
	}
	sort.Sort(elements)
}
