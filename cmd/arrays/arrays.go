package arrays

func Sum(arr []int) (sum int) {
	for _, val := range arr {
		sum += val
	}
	return
}

func SumAll(arr ...[]int) []int {
	lenarr := len(arr)
	sums := make([]int, lenarr)

	for i, num := range arr {
		println(num)
		sums[i] = Sum(num)
	}
	return sums
}
