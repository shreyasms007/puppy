// package main

import (
	"fmt"
	"math/rand"
)

func init() {

	fmt.Println("Begin initialization")
}

func dog1() string {

	return "is bigger version of puppy"
}

// func main() {
	

	s1 := dog1()
	compare(86)

	fmt.Println(s1)

	random(36)

}

func compare(x int) {

	if x < 42 {
		fmt.Println("Can print from 1 to 42")
	} else if x == 42 {
		fmt.Println("Number is exactly 42")
	} else {
		fmt.Println("Can print number more than 42")
	}

}

func random(y int) {

	if z := y * rand.Intn(y); z > y {
		fmt.Printf("Value of z and y are %v , %v\n and z is greater than y ", z, y)
	} else {
		fmt.Println("Value of z and y are %v , %v\n and  y is greater number", z, y)
	}
}



