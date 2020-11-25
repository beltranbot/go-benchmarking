package sort

import "sort"

// BubbleSort func
func BubbleSort(elements []int) {
	keepWorking := true
	for keepWorking {
		keepWorking = false
		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				keepWorking = true
				elements[i], elements[i+1] = elements[i+1], elements[i]
			}
		}
	}
}

// Sort func
func Sort(elements []int) {
	sort.Ints(elements)
}

// GetElements func
func GetElements(n int) []int {
	result := make([]int, n)
	for i := n - 1; i > 0; i-- {
		result[i] = i
	}

	return result
}
