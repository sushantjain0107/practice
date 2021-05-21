package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(nameQ string) *person {
	per := person{
		name: nameQ,
		age:  15,
	}

	//var perVar *person
	//perVar = &per
	//return perVar

	return &per
}

func main() {
	fmt.Println(person{"Bob", 20})

	p1 := person{
		name: "sushant",
		age:  56,
	}
	fmt.Println(p1)

	p2 := person{
		name: "Tashu",
	}
	fmt.Println(p2)
	fmt.Println("************==================================*************")
	//Using variable p1

	p11 := p1
	p1.age = 1000
	fmt.Println("p11 age is ", p11.age)
	fmt.Println("p1 age is ", p1.age)

	fmt.Println("&  * is ++++++++++++++++++  ", *&p1.age)

	fmt.Println("*************************")

	p22 := &p1
	p1.name = "Shanky"
	fmt.Println("p22 name is ", p22.name)
	fmt.Println("p1 name is ", p1.name)
	fmt.Println("*************************")

}
