package main

import "fmt"

func BubbleSort(array []int) {
	done := false
	n := len(array)

	for !done {
		done = true
		for i := 0; i < n-1; i++ {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
				done = false
			}
		}
	}
}

func InsertionSort(array []int) {
	n := len(array)
	for i := 1; i < n; i++ {
		for j := i; j >= 1 && array[j] < array[j-1]; j-- {
			array[j], array[j-1] = array[j-1], array[j]
		}
	}
}

func SelectionSort(array []int) {
	n := len(array)
	for i := 0; i < n-1; i++ {
		min_i := i
		for j := i + 1; j < n; j++ {
			if array[j] < array[min_i] {
				min_i = j
			}
		}

		if min_i != i {
			array[i], array[min_i] = array[min_i], array[i]
		}
	}
}

func Merge(array, left, right []int) {
	var i, j, k int

	// Merging
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			array[k] = left[i]
			i++
		} else {
			array[k] = right[j]
			j++
		}
		k++
	}

	// Copy what's left
	for i < len(left) {
		array[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		array[k] = right[j]
		j++
		k++
	}

}

func MergeSort(array []int) {
	if len(array) <= 1 {
		return
	}

	var q int = len(array) / 2

	left := make([]int, len(array[:q]))
	copy(left, array[:q])

	right := make([]int, len(array[q:]))
	copy(right, array[q:])

	MergeSort(left)  // left
	MergeSort(right) // right

	Merge(array, left, right)
}

func main() {
	n := []int{23, 4, 67, 9, 1221, 4, 80, 65, 33, 648, 90, 67, 8, 17, 64, 92, 11, 51, 172}

	fmt.Println("Original:", n)

	MergeSort(n)

	fmt.Println("Sorted:  ", n)
}
