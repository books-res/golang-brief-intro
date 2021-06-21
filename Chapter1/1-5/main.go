package main

import "fmt"

var msg string

func Hello(sb string) {
    fmt.Println("你好" + sb + "，" + msg)
}

func main() {
    msg = "见到你非常高兴"
    Hello("小明")
}
