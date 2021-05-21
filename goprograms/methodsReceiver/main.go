package main

import "fmt"

type rectangle struct {
	length int
	width  int
}

func (r rectangle) areaValue() int {
	return r.length * r.width
}

func (r *rectangle) areaPtr() int {
	return r.length * r.width
}

func main() {
	r := rectangle{
		length: 10,
		width:  5,
	}

	fmt.Println("Area through Value = ", r.areaValue())
	rPtr := &r
	fmt.Println("Area through Pointer = ", rPtr.areaPtr())

	r.length = 100
	fmt.Println("Area through Value = ", r.areaValue())
	fmt.Println("Area through Pointer = ", rPtr.areaPtr())

	r_1 := rectangle{
		length: 3,
		width:  2,
	}

	fmt.Println("%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%")
	r1 := r_1

	fmt.Println("Area through Value = ", r1.areaValue())
	fmt.Println("Area through Pointer = ", r1.areaPtr())
	r1.length = 8
	fmt.Println("Area through Value = ", r_1.areaValue())
	fmt.Println("Area through Poiner = ", r_1.areaPtr())
}
