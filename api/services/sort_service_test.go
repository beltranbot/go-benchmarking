package services

import (
	"testing"

	"github.com/beltranbot/go-benchmarking/api/utils/sort"
)

// integration testing
func TestSortLessThan10000(t *testing.T) {
	elements := sort.GetElements(10)
	Sort(elements)

	if elements[0] != 0 {
		t.Error("first element should be 0")
	}

	if elements[len(elements)-1] != 9 {
		t.Error("last elemetn should be 9")
	}
}

func TestSortMoreThan10000(t *testing.T) {
	elements := sort.GetElements(10001)
	Sort(elements)

	if elements[0] != 0 {
		t.Error("first element should be 0")
	}

	if elements[len(elements)-1] != 10000 {
		t.Error("last elemetn should be 10000")
	}
}

func BenchmarkBubbleSort10k(b *testing.B) {
	elements := sort.GetElements(10001)

	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}

func BenchmarkBubbleSort100K(b *testing.B) {
	elements := sort.GetElements(100001)

	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
