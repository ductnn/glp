package glp

func Abs(a int) int {
	if a > 0 {
		return a
	}

	return -a
}

func Swap(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}
