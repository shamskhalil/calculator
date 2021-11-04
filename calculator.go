package calculator

func Add(x *int, y *int) int {
	return *x + *y
}

func AddMulti(x ...int) int {
	sum := 0
	for _, val := range x {
		sum += val
	}
	return sum
}
