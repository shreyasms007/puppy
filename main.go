package main

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

func main() {

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


go func() {

	d1 := 1

	time.Sleep(d1*time.Millisecond)
	ch1 <- 41
}()

go func() {

	d2 := 1
	time.Sleep(d2*time.Millisecond)
	ch2 <- 42
}()

select {
case v1 := <-ch1 :
	fmt.Println("Print the value ",v1)
case v2 := <-ch2 :
	fmt.Println("Print the value ",v2)
}
