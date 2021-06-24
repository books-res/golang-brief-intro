package main

import "fmt"

func main() {
	var m map[string]int = map[string]int{"红烧肉": 88, "清蒸鱼": 98, "溜大虾": 128, "蒸螃蟹": 198, "鲍鱼粥": 68}
	var prices []*int
	for k, v := range m {
		// 本代码无法获取 m 的键的地址（&k 不是） 和 值的地址（&v不是，而&m[k]则提示语法错误：cannot take the address of m[k]）
		fmt.Printf("k:[%p].v:[%p]\n", &k, &v)
		//有问题的
		//	 fmt.Printf("k:[%p].v:[%p]\n", &k, &v)
		//	 prices = append(prices, &v)
		//正确的
		vv := m[k]
		prices = append(prices, &vv)
	}
	for _, price := range prices {
		fmt.Println(*price)
	}
}
