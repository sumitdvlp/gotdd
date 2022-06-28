package arrays

func Sum(arr []int) (sum int) {
	for _, val := range arr {
		sum += val
	}
	return
}

func SumAll(arr ...[]int) []int {
	var sums []int

	for _, num := range arr {
		sums = append(sums, Sum(num))
	}

	return sums
}

func SumAllTails(arr ...[]int) []int {
	var sums []int

	for _, num := range arr {
		trail := num[1:]
		sums = append(sums, Sum(trail))
	}

	return sums
}
