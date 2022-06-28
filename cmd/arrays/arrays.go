package arrays

func Sum(arr [5]int) (sum int) {
	for _, val := range arr {
		sum += val
	}
	return
}
