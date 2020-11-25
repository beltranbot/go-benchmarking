package sort

import (
	"testing"
)

func TestBubbleSortIncreaseOrder(t *testing.T) {
	elements := getElements(10001)
	BubbleSort(elements)

	if elements[0] != 0 {
		t.Error("first element should be 0")
	}

	if elements[len(elements)-1] != 10000 {
		t.Error("last element should be 10000", elements[len(elements)-1])
	}
}

func TestSortIncreaseOrder(t *testing.T) {
	elements := getElements(10001)
	Sort(elements)

	if elements[0] != 0 {
		t.Error("first element should be 0")
	}

	if elements[len(elements)-1] != 10000 {
		t.Error("last elemetn should be 9")
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	elements := GetElements(100000)

	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkSort(b *testing.B) {
	elements := GetElements(100000)

	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
