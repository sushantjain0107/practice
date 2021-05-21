package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned Normally")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from f -    ", r)
		}
	}()

	fmt.Println("Calling G")
	g(0)
	fmt.Println("Returned normally from G") //This will not printed
}

func g(i int64) {
	if i > 2 {
		fmt.Println("Start Panicking ")
		panic(fmt.Sprintf("%v ", i))
	}
	defer fmt.Println("Defer in g ", i)

	fmt.Println("Printing in g ", i)
	g(i + 1)
}
