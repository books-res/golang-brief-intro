package main

import "fmt"

func main() {
	var foods = []string{"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "鲍鱼粥"}
	var other_foods = foods[:]
	fmt.Println(other_foods)

	var foodsSlice2 = []string{"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "鲍鱼粥"}
	foodsSlice6 := []string{"米饭", "面条"}
	copy(foodsSlice6, foodsSlice2)
	fmt.Println(foodsSlice2)
	fmt.Println(foodsSlice6)

	var foodsSlice3 = []string{"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "鲍鱼粥"}
	foodsSlice7 := []string{"米饭", "面条"}
	copy(foodsSlice3, foodsSlice7)
	fmt.Println(foodsSlice3)
	fmt.Println(foodsSlice7)

	// 新添加的实验代码
	var a1 = [8]string{"a", "b", "c", "d", "e", "f", "g", "h"}
	var a2 = [...]string{"A","B","C"}
	fmt.Printf("a1 =\t%s\n", a1)
	fmt.Printf("a2 =\t%s\n", a2)
	// 1 从切片种提取切片
	var a11 = a1[1:]
	fmt.Printf("%s\n", a11)
	// 2 copy(dest, src []Type) int
	a21 := a2[1:]
	a12 := a1[2:]
	copy(a21, a12)
	fmt.Printf("%s <- %s\n", a21, a12)

}
