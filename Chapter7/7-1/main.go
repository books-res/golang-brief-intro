package main

import "fmt"

// 无返回值函数。这里是函数的注释部分
func f1(x, y int) {}

// 有1个返回值, 2个返回值， 多个返回值
func f2(x, y int) int {
	return 1
}
func f3(x, y int) (result int) {
	result = 1
	return
}
func f4(x, y int) (int, int) {
	return 1,2
}
func f5(x int, y int) (int, int, int) {
	return 1,2,3
}

func useTestFunc() {
	f1(1,2)
	r2 := f2(1,2)
	fmt.Println("r2 =", r2)
	r3 := f3(1,2)
	fmt.Println("r3 =", r3)
	r41, r42 := f4(1,2)
	fmt.Println("r4 =", r41, r42)
	f5(1,2)
}

//方法注释
func Fav() {
	fmt.Println("我喜欢吃生蚝")
}

func add(x, y int) (result int) {
	result = x + y
	return
}

func main() {
	useTestFunc()
	Fav()
}
