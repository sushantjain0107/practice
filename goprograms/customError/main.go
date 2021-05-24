package main

import (
	"errors"
	"fmt"
)

func main() {
	side := -2

	area, error := findArea(side)

	if error != nil {
		fmt.Printf("%s", error)
		return
	}
	fmt.Printf("Area is %d ", area)

}

func findArea(side int) (int, error) {

	if side < 0 {
		return 0, errors.New("negative value")
	}

	return side * side, nil
}
