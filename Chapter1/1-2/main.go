package main

import (
    "fmt"
)

// 这是一个单行注释

/*
 这是一个多行注释
*/

func main() {
    var str string = "一个字符串"
    num := 123
    var strPtr = &str
    fmt.Println("str = ", str)
    fmt.Println("num = ", num)
    fmt.Println("strptr = ", strPtr)
    fmt.Println("*strptr = ", *strPtr)
}
