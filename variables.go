package main

import (
	"fmt"
	"math/rand"
)

func main() {

	x := rand.Intn(10)
	y := rand.Intn(10)

	fmt.Println("Value of x is :", x)
	fmt.Println("Value of y is :", y)

	if x < 4 && y < 4 {

		fmt.Println("both are less than 4")
	} else if x > 6 && y > 6 {

		fmt.Println("Both are greater than 6")
	} else if x >= 4 && x <= 6 {
		fmt.Println("Value of x is :", x)
	} else if y != 5 {
		fmt.Println("value of y is :", y)
	} else if y == 5 {
		fmt.Println("Value of x is : ", x, "and Value of y is :", y)
	} else {
		fmt.Println("shreyas.")
	}

}
