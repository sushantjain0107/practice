package main

import "fmt"

func main() {
	map1 := make(map[string]string)

	map1["name"] = "Sushant"
	map1["age"] = "30"

	delete(map1, "qwe")

	value, isPresent := map1["name"]
	fmt.Println("key is ", value)
	fmt.Println("val is ", isPresent)

	map2 := map[string]string{"abc": "aa", "xyz": "xx", "qwer": "qqq"}

	for i, j := range map2 {
		fmt.Print(i, "-")
		fmt.Println(j)
	}
}
