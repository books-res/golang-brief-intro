package main

import "fmt"

func main() {
	var m map[string]int = map[string]int{"红烧肉": 88, "清蒸鱼": 98, "溜大虾": 128, "蒸螃蟹": 198, "鲍鱼粥": 68}
	price, ok := m["鲍鱼粥"]
	if !ok {
		fmt.Println("没有对应的键值对")
	} else {
		fmt.Println(price)
	}

	// 这里可以简单的认为 interface{} 为任意类型，所以键可以是string和int。但实际不是任意类型，具体学了interface之后就明白了。
	var errorMap = map[interface{}]int{
		 "红烧肉": 88,
		 123: 112233,
		//"红烧肉": 89, //重复的key提示语法错误: duplicate key "红烧肉" in map literal
		//[]string{"清蒸鱼"}: 98, // panic: runtime error: hash of unhashable type []string
		"溜大虾": 128,
	}
	fmt.Println(errorMap)
}
