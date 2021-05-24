package main

import (
	"fmt"
	"time"
)

func main() {
	//go numbers()
	go alphabet()

	time.Sleep(5 * time.Second)
	fmt.Println("Makn terminated ")
}

func alphabet() {

	for i := 'a'; i <= 'e'; i++ {
		fmt.Printf("%c \n", i)
	}
}
