package app

func SumArray(arr [5]int) int {
	var sum int
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
	}
	return sum
}