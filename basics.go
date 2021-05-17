package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	array()
	array2D()
	arraySlice()

	/*Data types - Explicit*/
	var number uint8 = 255
	var number1 uint16 = 266
	fmt.Println("Hello World!", number, number1)

	/*Data types - Implicit*/
	var num = 200.98
	fmt.Printf("Type is: %T \n", num)

	/* Assignment Expression */
	aeOperator := 6 //short variable declarations
	fmt.Printf("Type is: %T \n", aeOperator)

	/* Printing to console and fmt*/
	fmt.Printf("check type and value: %T %v \n", 10, 10)
	fmt.Printf("checkboolean:  %t \n", true)
	fmt.Printf("Binary Representation: %b \n", 1025)
	fmt.Printf("Octal Representation: %o \n", 1025)
	fmt.Printf("Decimal Representation: %d \n", 1025)
	fmt.Printf("String Representation: %s \n", "sapna")
	fmt.Printf("Double String Representation: %q \n", "sapna")

	/* Console Input(Bufio Scanner & Type Conversion)*/
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type the year you were born: ")
	scanner.Scan()
	input, err := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("value is: %d \n", input)
	fmt.Printf("ERR is: %v \n", err)

	/* Arithmetic Operators & Math*/
	var num1 float64 = 8
	var num2 int = 5
	sum := num1 / float64(num2)
	fmt.Printf("%f \n", sum)

	/*Conditions & Boolean Expression */
	x := 5
	y := 6.5
	val := float64(x)+1.5 != y
	fmt.Printf("%t \n", val)

	a := "sap"
	b := "Sap"
	compare := a < b
	fmt.Printf("%t \n", compare)

	/*Chained Conditionals AND, OR, NOT */
	value := (true || false) && !!false
	value1 := value || false
	fmt.Printf("%t \n", value1)

	/**If, Else If, Else */
	age := 19
	fmt.Println("Before If")
	if age >= 18 {
		fmt.Println("Eligible")
	} else if age >= 14 {
		fmt.Println("Not Eligible")
		fmt.Printf("wait %d years \n", 18-age)
	} else {
		fmt.Println("Not Eligible")
	}

	/*For Loop and While Loop*/
	u := 4
	for u < 7 {
		fmt.Println(u)
		u++
	}

	for a := 0; a <= 5; a++ {
		// fmt.Println(a)
		// break
		continue
	}

	/*Switch statement*/
	ans := 1
	switch ans {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("No case")
	}

	switch {
	case ans > 0:
		fmt.Println("greater then 0")
	case ans < 0:
		fmt.Println("less then 0")
	default:
		fmt.Println("0")
	}

}

func array() {
	var arr [5]int
	arr[0] = 100
	fmt.Println("Array", arr)

	newArr := [3]int{4, 5, 6}
	sum := 0
	for i := 0; i < len(newArr); i++ {
		sum += newArr[i]
	}

	fmt.Println("sum", sum)
}

func array2D() {
	arr2D := [2][3]int{{1, 2, 3}, {3, 3, 4}}
	fmt.Println("array2D", arr2D[1][2])
}

func arraySlice() {
	var x [5]int = [5]int{1, 2, 3, 4, 5}
	var s []int = x[1:3]
	fmt.Println("slice", s)
	fmt.Println("length", len(s))
	fmt.Println("capacity", cap(s))

	var a []int = []int{1, 2, 3, 4, 5}
	b := append(a, 10)
	fmt.Println("append", b)

}
