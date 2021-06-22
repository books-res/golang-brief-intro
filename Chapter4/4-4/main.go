package main

import "fmt"

func main() {
	foods := []string{"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "鲍鱼粥"}
	foods2 := make([]*string, len(foods))
	// 错误的示例
	//for i, value := range foods {
	//	foods2[i] = &value
	//}
	// 正确的示例
	for i, _ := range foods {
		foods2[i] = &foods[i]
	}
	fmt.Println(foods[0], foods[1], foods[2], foods[3], foods[4])
	fmt.Println(*foods2[0], *foods2[1], *foods2[2], *foods2[3], *foods2[4])
}
