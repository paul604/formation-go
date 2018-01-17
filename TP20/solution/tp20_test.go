package sorting

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
)

var biglist = []int{}

func TestMain(m *testing.M) {
	fmt.Println("Init")
	for i := 0; i < 1000; i++ {
		biglist = append(biglist, rand.Intn(100))
	}
	os.Exit(m.Run())
}

func BenchmarkBubbleSortbiglist(b *testing.B) {
	list := make([]int, 1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(list, biglist)
		bubblesort(list)
	}
}

func BenchmarkSelctionSortbiglist(b *testing.B) {
	list := make([]int, 1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(list, biglist)
		selectionSort(list)
	}
}

func BenchmarkQuickSortbiglist(b *testing.B) {
	list := make([]int, 1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(list, biglist)
		quicksort(list)
	}
}
