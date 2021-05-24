package basics

import "fmt"

var gVar float64 //Global Variable

func LocalGlobal() {
	var lVar1 float64 //Local variable
	lVar2 := 20.0
	lVar1 = 30.5

	gVar = lVar1 * lVar2

	fmt.Printf("Sum of %f and %f is %f ", lVar1, lVar2, gVar)
}
