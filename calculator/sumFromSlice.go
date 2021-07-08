package calculator

func SumFromSlice(data []int) int {
	sumValue := 0
	for _, val := range data {
		sumValue += val
	}
	return sumValue
}
