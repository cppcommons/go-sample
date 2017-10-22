package main

import (
	"fmt"
)

var sum func(i, j int) = func(i, j int) {
	fmt.Println(i + j)
}

func sum2(i, j int) {
	fmt.Println(i + j)
}

func main() {
	x := 123
	var y int32
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println("hello world")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	var j int64
	for j = 0; j < 10; j++ {
		fmt.Println(j)
	}
	n := 10
	switch n {
	case 15:
		fmt.Println("FizzBuzz")
	case 5, 10:
		fmt.Println("Buzz")
	case 3, 6, 9:
		fmt.Println("Fizz")
	default:
		fmt.Println(n)
	}
	func(i, j int) {
		fmt.Println(i + j)
	}(2, 4)
	locFun := func(i, j int) {
		fmt.Println(i + j)
	}
	locFun(11, 22)
	sum(22, 33)
	sum2(22, 33)
}
