package basic

import (
	"fmt"
)

var hra int = 5
var tax int = 2
var Basic int

func Calculation() (allowance int, deduction int) {
	fmt.Println("Inside basic")
	allowance = (Basic * hra) / 100
	fmt.Println("allowance", allowance)
	deduction = (Basic * tax) / 100
	fmt.Println("deduction", deduction)
	return
}
