package glp

// Bubble sort
func BubbleSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-1-i; j++ {
			if array[j] > array[j+1] {
				Swap(array, j, j+1)
			}
		}
	}

	return array
}

// Insertion sort
func InsertionSort(array []int) []int {
	for i := 1; i < len(array); i++ {
		for j := i; j > 0 && array[j] < array[j-1]; j-- {
			Swap(array, j, j-1)
		}
	}

	return array
}

// Selection sort
func SelectionSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[i] {
				Swap(array, i, j)
			}
		}
	}
	return array
}

// Quick sort
func QuickSort(array []int) []int {
	Helper(array, 0, len(array)-1)

	return array
}
