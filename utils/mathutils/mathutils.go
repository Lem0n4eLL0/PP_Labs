package mathutils

func Factorial(num int64) int64 {
	if num == 0 {
		return 1
	}
	return Factorial(num-1) * num
}
