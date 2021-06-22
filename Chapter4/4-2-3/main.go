package main

import "fmt"

func main() {
	// 从切片种删除所为3的元素
	var foods = []string{"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "鲍鱼粥"}
	fmt.Println("before: ", foods)
	foods2 := append(foods[:3], foods[4:]...)
	fmt.Println("  after:", foods2)
}
