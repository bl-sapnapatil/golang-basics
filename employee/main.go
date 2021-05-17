package main

import (
	b "employee/basic"
	"employee/basic/gross"
	"fmt"
)

func main() {
	fmt.Println("Hii")
	b.Basic = 10000
	fmt.Println("salary", gross.GrossSalary())
}
