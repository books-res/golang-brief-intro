package main

import (
    "fmt"
    "github.com/golang-brief-intro/myutils"
    "strings"
)

func main() {
	foods3 := "    米饭，馒头，面条    "
	foods3 = strings.Trim(foods3, " ")
	fmt.Println(foods3)

	foods4 := "米饭，馒头，面条"
	result := strings.Trim(foods4, "米饭")
	fmt.Println(result)

    str1 := "   1 2  3   4    5     1 2  3   4    5     "
    myutils.PrintlnLR(str1)                   // 原样输出
    myutils.PrintlnLR(strings.Trim(str1,""))  // 空字符串
    myutils.PrintlnLR(strings.Trim(str1," ")) // 去掉左右两侧空格

    str2 := "go hello go world    "
    myutils.PrintlnLR(str2)
    myutils.PrintlnLR(strings.Trim(str2,"go")) // 去掉左侧部分

    str3 := "go hello go world ~! gogogo"
    myutils.PrintlnLR(str3)
    myutils.PrintlnLR(strings.Trim(str3, "go")) // 去掉两侧
}
