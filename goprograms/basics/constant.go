package basics

import "fmt"

func ConstantDemo() {
	const LENGTH float64 = 10.3
	const WIDTH float64 = 10.5
	var area float64 = LENGTH * WIDTH
	fmt.Println(" \"Area is \" ", area)
	fmt.Println("Area is using %d ", area)
	fmt.Printf(" \"Area izzzz perc T %T \" \n", area)
	fmt.Printf(" \"Area izzzz perc F %F \" \n", area)
	fmt.Printf(" \"Area izzzz perc d %d \" \n\n", area)

	const L int = 2
	const W int = 3
	var are int = L * W
	fmt.Printf("ARE is : %d  ", are)
}
