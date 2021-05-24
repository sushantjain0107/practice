package basics

import "fmt"

func VariadicDemo(name string, numb ...int) {

	for index, value := range numb {
		fmt.Printf("Index %d , Value %d \n", index, value)
	}
}
