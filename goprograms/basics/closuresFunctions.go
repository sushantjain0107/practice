package basics

//Ananymous function
func IntSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func TwoTable() func() int {
	i := 0
	return func() int {
		i = i + 2
		return i
	}
}
