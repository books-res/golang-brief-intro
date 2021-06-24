package main

func Desc(foodType string, foodName string) {
	//函数形参在函数内可以不用
	//fmt.Println("这是" + foodType)
	//fmt.Println("这是" + foodType + "的" + foodName)
}

func main() {
	// 函数如有形参，调用时必须提供实参
	//not enough arguments in call to Desc
	//Desc()
	Desc("食物类型","食物名称")
}
