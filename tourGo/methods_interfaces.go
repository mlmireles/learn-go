package main

import (
	"fmt"
	"math"
)

type Mertex struct {
	X,Y float64
}

/* 
	Go does not have classes. However, you can define methods on types.
	A method is a function with a special 'receiver' argument.
	The receiver appears in its own argument list between the func keyword
	and the method name
*/
func(v Mertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func methods() {
	v := Mertex{3, 4}
	fmt.Println(v.Abs())
}

func methodsInterfaces() {
	methods()	
}
