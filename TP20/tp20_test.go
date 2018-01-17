package main

import (
	"testing"
)

var biglist = []int{}

func BenchmarkBubbleSort(b *testing.B) {
	list := make([]int, 1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(list, biglist)
		bubblesort(list)
	}
}

func BenchmarkSelctionSort(b *testing.B) {
	list := make([]int, 1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(list, biglist)
		selectionSort(list)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	list := make([]int, 1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(list, biglist)
		quicksort(list)
	}
}
