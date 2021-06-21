package main

import "fmt"

func main() {
	var list = [5]string{
		"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "蒜蓉粉丝扇贝",
	}
	println("第1次输出：使用 %")
	for idx, item := range list {
		desc := fmt.Sprintf("%d-%s", idx, item)
		fmt.Println(desc)
	}
	println("第2次输出：使用#{}")
	for idx, item := range list {
		desc := fmt.Sprintf("%d-%s", idx, item)
		fmt.Println(desc)
	}
}
