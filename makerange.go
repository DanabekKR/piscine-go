package piscine

func MakeRange(min, max int) []int {
	var emp []int
	if min >= max {
		return emp
	}
	res := make([]int, max-min)
	for i := min; i < max; i++ {
		res[i-min] = i
	}
	return res
}
