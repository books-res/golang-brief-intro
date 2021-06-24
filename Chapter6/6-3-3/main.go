package main

import "fmt"

func main() {
	// 修改
	var m map[string]int = map[string]int{"红烧肉":88,"清蒸鱼":98, "溜大虾":128, "蒸螃蟹":198, "鲍鱼粥":68}
	m["蒸螃蟹"]=398
	fmt.Println(m)
}
