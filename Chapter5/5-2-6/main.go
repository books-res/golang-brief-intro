package main

import "fmt"

func main() {
    choosePrice := []int{1000, 1288, 1688, 2000}
    for _, price := range choosePrice {
        fmt.Print("price=", price, " => ")
        if price < 1288 {
            fmt.Println("没有优惠")
        } else if price == 1288 {
            fmt.Println("赠送优惠卡一张")
        } else if price == 1688 {
            fmt.Println("赠送优惠卡一张 + 赠送果盘一个")
        } else {
            fmt.Println("赠送优惠卡两张")
        }
    }
}
