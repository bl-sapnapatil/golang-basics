package gross

import (
	b "employee/basic"
	"fmt"
)

func GrossSalary() int {
	fmt.Println("Inside Gross");
	a, t := b.Calculation()
	fmt.Println("value of a and t", a, t)
	return ((b.Basic + a) - t)
}
