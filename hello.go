package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

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


}
