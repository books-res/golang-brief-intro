package main

import "fmt"

var msg = "包内私有全局变量"

func Cook() {
    msg := "Cook函数内局部变量"
    println(msg)
}

func main() {
    fmt.Println(msg)

    msg := "main函数内局部变量"
    fmt.Println(msg)

    Cook()
    fmt.Println(msg)

    {
        msg := "代码块内局部变量"
        fmt.Println(msg)
    }
    fmt.Println(msg)
}
