package main

import "fmt"

func main() {
	var foods = []string{"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "鲍鱼粥"}
	fmt.Println(foods)
	foods = append(foods, "三文鱼")
	fmt.Println(foods)

	// 切片后插入元素
	println()
	foods2 := []string{}
	foods2 = append(foods2, "红烧肉")
	foods2 = append(foods2, "清蒸鱼")
	foods2 = append(foods2, "溜大虾", "蒸螃蟹", "鲍鱼粥")
	fmt.Println(foods2)

	// 扩展5个元素长度
	println()
	foods3 := []string{"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "鲍鱼粥"}
	fmt.Println("before:", foods3)
	foods3 = append(foods3, make([]string, 5)...)
	fmt.Println("after: ", foods3)
	fmt.Printf("长度：%d\n", len(foods3))
	fmt.Printf("容量：%d\n", cap(foods3))

	// 索引2的位置插入1个元素
	println()
	foods4 := []string{"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "鲍鱼粥"}
	foods5 := append(foods4[:2], append([]string{"三文鱼"}, foods4[2:]...)...)
	fmt.Println(foods5)

	// 索引4的位置插入长度为3的新切片（例子是长度为3的空切片）
	println()
	foods6 := []string{"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "鲍鱼粥"}
	foods7 := append(foods6[:4], append(make([]string, 3), foods6[4:]...)...)
	fmt.Println(foods7)

	//索引4的位置插入新切片
	println()
	foods8 := []string{"米饭", "面条", "馒头"}
	var foods9 = []string{"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "鲍鱼粥"}
	foods10 := append(foods9[:4], append(foods8, foods9[4:]...)...)
	fmt.Println(foods10)
}
