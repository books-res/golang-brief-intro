package main

import "fmt"

func main() {
	var foods = [5]string{"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "鲍鱼粥"}
	var foodsSlice []string = foods[0:3]
	fmt.Println(foodsSlice)
	fmt.Println(foods[1:4])
	fmt.Println(foods[:2])
	fmt.Println(foods[3:])

	var foodsSlice2 = []string{"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "鲍鱼粥"}

	foodsSlice5 := foodsSlice2[1:4]
	fmt.Println(len(foodsSlice5))
	fmt.Println(cap(foodsSlice5))

	println("\n演示切片扩容")
	foodSlice4 := []string{"红烧肉"}
	fmt.Println(foodSlice4)
	fmt.Print(fmt.Sprintf("Len:%d\t", len(foodSlice4)))
	fmt.Println(fmt.Sprintf("Cap:%d", cap(foodSlice4)))
	foodSlice4 = append(foodSlice4, "清蒸鱼")
	fmt.Println(foodSlice4)
	fmt.Print(fmt.Sprintf("Len:%d\t", len(foodSlice4)))
	fmt.Println(fmt.Sprintf("Cap:%d", cap(foodSlice4)))
	foodSlice4 = append(foodSlice4, "溜大虾")
	fmt.Println(foodSlice4)
	fmt.Print(fmt.Sprintf("Len:%d\t", len(foodSlice4)))
	fmt.Println(fmt.Sprintf("Cap:%d", cap(foodSlice4)))
	foodSlice4 = append(foodSlice4, "蒸螃蟹")
	fmt.Println(foodSlice4)
	fmt.Print(fmt.Sprintf("Len:%d\t", len(foodSlice4)))
	fmt.Println(fmt.Sprintf("Cap:%d", cap(foodSlice4)))
	foodSlice4 = append(foodSlice4, "鲍鱼粥")
	fmt.Println(foodSlice4)
	fmt.Print(fmt.Sprintf("Len:%d\t", len(foodSlice4)))
	fmt.Println(fmt.Sprintf("Cap:%d", cap(foodSlice4)))

	// 新添加的实验代码
	var a1 = [8]string{"a", "b", "c", "d", "e", "f", "g", "h"}
	var a2 = [...]string{"A","B","C"}
	fmt.Printf("a1 =\t%s\n", a1)
	fmt.Printf("a2 =\t%s\n", a2)
	// 没有python那么强大的切片功能
	// 1.基于数组创建切片
	fmt.Printf("a1[2:6] =\t%s\n", a1[2:6])
	fmt.Printf("a1[:6] =\t%s\n", a1[:6])
	fmt.Printf("a1[2:] =\t%s\n", a1[2:])
	fmt.Printf("a1[:] =\t\t%s\n", a1[:])
	// 2.直接创建切片
	var s1 = []string{"A","B","C"}
	fmt.Printf("s1=%s\n", s1)
	// 3.函数make([] type, len, <cap>)创建切片
	var s2 = make([]string, 3, 5)
	fmt.Printf("s2=%s\tlen(s2)=%d\tcap(s2)=%d\n", s2, len(s2), cap(s2))
	for i, e := range s2{
		fmt.Printf("%d -> [%s]\n", i, e)
	}
	s2[2], s2[1] = "-2", "-1"
	fmt.Printf("s2=%05s\tlen(s2)=%d\tcap(s2)=%d\n", s2, len(s2), cap(s2))
}
