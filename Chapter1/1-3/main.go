package main

import "fmt"

func main() {
    str := "hello go world"
    var strPtr = &str
    fmt.Println("str = ", str)
    fmt.Println("strPtr = ", strPtr)
    fmt.Println("*strPtr = ", *strPtr)

    println()

    var strPtr2 = new(string)
    fmt.Println("strPtr2 = ", strPtr2)
    fmt.Println("*strPtr2 = [", *strPtr2, "]")
    *strPtr2 = "Hello Go ~!"
    println("strPtr2 = ", strPtr2)
    println("*strPtr2 = [", *strPtr2, "]")
}
