package basics

func Recursion(numb int64) int64 {
	if numb == 0 {
		return 1
	}
	return numb * Recursion(numb-1)
}
