package main

import "fmt"

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


}
