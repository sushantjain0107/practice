package main

import (
	"fmt"
	"strings"
)

// Main function
func main_split() {

	// Creating and initializing the strings
	str1 := "Welcome, to the, online portal, of GeeksforGeeks"
	str2 := "My dog name is Dollar"
	str3 := "I like to play Ludo"

	// Displaying strings
	fmt.Println("String 1: ", str1)
	fmt.Println("String 2: ", str2)
	fmt.Println("String 3: ", str3)

	// Splitting the given strings
	// Using Split() function
	res1 := strings.Split(str1, ",")
	res2 := strings.Split(str2, "")
	res3 := strings.Split(str3, "!")
	res4 := strings.Split("", "GeeksforGeeks, geeks")

	fmt.Println("++++++++++++++++", len(res1))
	// Displaying the result

	fmt.Println("\nResult 1: ", res1)
	fmt.Println("Result 2: ", res2)
	fmt.Println("Result 3: ", res3)
	fmt.Println("Result 4: ", res4)

}

func main() {
	// Creating and initializing the strings
	str1 := "Welcome, to the, online portal, of GeeksforGeeks"
	str2 := "My dog name is Dollar"
	str3 := "I like to play Ludo"

	fmt.Println(str1[1:4])
	fmt.Println(str2[0:7])
	fmt.Println(str3[5:15])
}
