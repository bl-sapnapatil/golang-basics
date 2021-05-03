package main

import "fmt"

func main() {

    /*Data types - Explicit*/
    var number uint8 = 255;
    var number1 uint16 = 266;
    fmt.Println("Hello World!",number,number1)

    /*Data types - Implicit*/
    var num = 200.98
    fmt.Printf("%T",num)

    /* Assignment Expression */
    aeOperator := 6; //short variable declarations
    fmt.Printf("%T",aeOperator)

}