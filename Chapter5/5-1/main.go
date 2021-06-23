package main

import "fmt"

func main() {
	foods := []string{"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "鲍鱼粥"}

	// 简单遍历
	for i, e := range foods {
		fmt.Println(i, "->", e)
	}

	for _, item := range foods {
		if item == "鲍鱼粥" {
			fmt.Println("鲍鱼粥，今日免费赠送")
		} else {
			fmt.Println(item)
		}
	}
}