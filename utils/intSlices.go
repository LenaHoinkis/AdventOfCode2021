package utils

func SumInts(data []int) int {
	sum := 0
	for _, value := range data {
		sum += value
	}
	return sum
}

func ProdInts(data []int) int {
	sum := 1
	for _, value := range data {
		sum *= value
	}
	return sum
}

func MinInts(data []int) int {
	min := data[0]
	for _, i := range data {
		if min > i {
			min = i
		}
	}
	return min
}

func MaxInts(data []int) int {
	max := data[0]
	for _, i := range data {
		if max < i {
			max = i
		}
	}
	return max
}

func Greater(x int, y int) int {
	if x > y {
		return 1
	}
	return 0
}

func Less(x int, y int) int {
	if x < y {
		return 1
	}
	return 0
}
func Equal(x int, y int) int {
	if x == y {
		return 1
	}
	return 0
}
