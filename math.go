package glp

// int
func AddInt(a int, b int) int {
	return a + b
}

func MultifyInt(a, b int) int {
	return a * b
}

func DevideInt(a, b int) int {
	return a / b
}

func SqrtInt(a int) int {
	return a * a
}

func IsEven(a int) bool {
	return a&1 == 0
}

func Abs(a int) int {
	if a > 0 {
		return a
	}

	return -a
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Swap(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}
