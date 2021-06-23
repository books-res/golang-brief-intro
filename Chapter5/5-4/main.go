package main

import "fmt"

func main() {
    total := 0

    //第1种for循环
    for i := 0; i < 10; i++ {
        total += i
    }

    total2 := 0

    //第2种for循环
    for {
        total2++
        if total2 > 100 {
            break
        }
    }

    total3 := 0
Out:
    for {
        total3++
        if total3 > 10 {
            // 跳出外层循环
            break Out
        }
    }

    foods := []string{"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "鲍鱼粥"}
    //第3种for循环
    for i := range foods {
        fmt.Println(i)
    }

    //foods2 := []string{"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "鲍鱼粥"}
    foods2 := []string{"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "鲍鱼粥"}
    //第4种for循环
    for i, item := range foods2 {
        if i == len(foods2) {
            item = "蒜蓉粉丝扇贝"
        }
        result := fmt.Sprintf("%d--%q", i, item)
        fmt.Println(result)
    }

    fmt.Println(foods2)

}
