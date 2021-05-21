package main

import (
	"fmt"
	abc "goprograms/basics"
)

func main_1() {
	fmt.Print("Hello \n")
	/*
		i := 0
		for i < 10 {
			i = i + 2
			fmt.Println("\"Count is \" ", i)
		}
		abc.ConstantDemo()
		abc.LocalGlobal()
	*/
	//abc.ConstantDemo()
	//abc.MapDemo()cls
	//abc.VariadicDemo("Sushant", 25)

	/*
		sequence := abc.IntSeq()
		fmt.Println(sequence())
		fmt.Println(sequence())
		fmt.Println(sequence())
		fmt.Println(sequence())

		sequence2 := abc.IntSeq()
		fmt.Println(sequence2())
		fmt.Println(sequence2())
	*/

	/*
		tabelOfTwo := abc.TwoTable()
		fmt.Println(tabelOfTwo())
		fmt.Println(tabelOfTwo())
		fmt.Println(tabelOfTwo())
		fmt.Println(tabelOfTwo())
		fmt.Println(tabelOfTwo())
	*/

	fmt.Println(abc.Recursion(7))
}

//For pointer
func main() {

	//abc.MapDemo()
	/*
		i := 1
		fmt.Println("Initial - ", i)

		zeroVal(i)
		fmt.Println("Zero Val - ", i)

		zeroPtr(&i)
		fmt.Println("Zero Ptr - ", i)

		fmt.Println("Print Pointer - ", &i)
	*/
}

/*
func zeroVal(iVar int) {
	iVar = 0
}

func zeroPtr(iPtr *int) {
	*iPtr = 0
}
*/
