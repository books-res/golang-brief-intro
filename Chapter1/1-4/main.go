package main

import (
    "fmt"
    // 导入自定义的包
    "github.com/golang-brief-intro/Chapter1/1-4/mypack"
)

func main() {
    // 访问自定义包导出的变量
    fmt.Println(mypack.Str)
    // 访问自定义包导出的函数
    mypack.SayHello1("Go")
    // 小写开头的变量或函数无法在包外访问
    // fmt.Println(mypack.str)
    // mypack.sayHello2("go")
}
