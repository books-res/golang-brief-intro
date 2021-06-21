package mypack

import "fmt"

// 包外可以访问，全局变量
var Str string

func init() {
    fmt.Println("print message from hello1.go")
    Str = "a string init from hello1.go init()"
}

func SayHello1(sb string) {
    fmt.Println("Hello", sb, "~!")
    // str只能在包内访问
    fmt.Println("str = ", str)
}
