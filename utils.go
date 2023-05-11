package glp

func Helper(array []int, left, right int) {
	if len(array) <= 1 {
		return
	}

	if left < right {
		pivot := Partition(array, left, right)
		Helper(array, left, pivot-1)
		Helper(array, pivot+1, right)
	}
}

func Partition(array []int, left, right int) int {
	index := left - 1
	// Defind pivot element
	p := array[right]

	for i := left; i < right; i++ {
		if array[i] <= p {
			index++
			Swap(array, index, i)
		}
	}

	Swap(array, index+1, right)

	return index + 1
}
