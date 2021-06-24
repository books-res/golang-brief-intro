package main

import "fmt"

// 数组形参，注意不是切片形参

func ShowFoods(foods [3]string) {
	//修改第一个元素
	foods[0] = "蛋炒饭"
	fmt.Println(foods)
}

//数组指针
func ShowFoods2(foods *[3]string) {
	foods[0] = "蛋炒饭"
	fmt.Println(foods)
}

func main() {
	todayFoods := [3]string{"卤肉饭", "西红柿鸡蛋饭", "肥牛饭"}
	ShowFoods(todayFoods)
	fmt.Println(todayFoods)

	todayFoods2 := &[3]string{"卤肉饭", "西红柿鸡蛋饭", "肥牛饭"}
	ShowFoods2(todayFoods2)
	fmt.Println(todayFoods2)

}
