package main

import ( 
	"fmt" 
	"math"
)

var c, python, java bool
var cpp, python_, java_ = true, false, "no!"

func add(x, y int) int { 
	return 	x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println("Now you have %g problems.", math.Sqrt(7))
	fmt.Println(add(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(19))
}
