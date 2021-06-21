// 一个文件夹内的go文件声明所属的包必须一样
//package mypack2
package mypack

import "fmt"

// 包外无法访问，全局变量
var str string

func init() {
    fmt.Println("print message from hello2.go")
    str = "a string cannot be require from other package"
}

// 包外无法访问
func sayHello2(sb string) {
    fmt.Println("Hello", sb, "~!")
}
