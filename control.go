package main

import (
    "fmt"
)


func if_control(){
    a := 10

    if a > 0 {              //表达式无括号, 左大括号
        a += 100
    } else if a == 0 {
        a = 0
    } else {
        a -= 100
    }

    if a > 0 { a += 100 } else { a -= 100 }     //单行模式

    if b := 100; b != 0{        //支持初始化表达式
        println(b)
    }
    println(a)

    c := 10
    rlt := map[bool]int{true: a, false: c}[a > c] //不支持?:，替代做法
    fmt.Printf("%T, %v\n", rlt, rlt)
}

/*
for支持三种形式：
    for init; condition; post {}
    for condition {}
    for {}
for ... range 用于完成迭代器操作， 可用于string array slice map channel
返回index:value key:value, map key, channel value
*/
func for_control() {
    ss := "abcd"
    for i, m := 0, len(ss); i < m; i++ {
        println(ss[i])
    }

    d := func() []int {
        println("data...")
        return []int{1, 2, 3}
    }
    for i, x := range d() {                     //迭代右表达式在循环前一次性计算，迭代变量每次循环时拷贝赋值
        fmt.Printf("%d: %d, %p\n", i, x, &x)
    }

    array := []int{1, 2, 3}
    for i, x := range array {
        fmt.Printf("array[%d] = %d\n", i, x)    //迭代变量总是从array复制中获得
        x++                                     //修改的是复制品，不影响array
        if i+1 < len(array) {
            array[i+1] += 1                     //修改的是原对象
        }
    }
    fmt.Println(array)

    slice := []int{1, 2, 3}
    for i, x := range slice {
        fmt.Printf("slice[%d] = %d\n", i, x)
        if i == 0 { slice[1] = 100}             //引用类型，共享底层数组，修改有效
        slice = append(slice, x+100)            //修改slice原对象，不影响复制品
    }
    fmt.Println(slice)

    maps := map[string]int{"x": 1, "y": 2, "z": 3} //引用类型的复制品, 操作会影响原对象
    for k, v := range maps {
        fmt.Printf("maps[%s] = %d\n", k, v)
        delete(maps, k)                             //安全删除元素项
    }       //由于map使用随机存储，没有固定迭代次序，不建议迭代时添加元素项，以免造成不可预知结果
    fmt.Println(maps)
}

func switch_control() {
    switch a := 5; a {
    case 0, 1:              //逗号指定多个分支
        println("a")
    case 100:               //什么都不做，不是fallthrough
    case 5:
        println("b")        //多行，不用{}
        fallthrough         //直接进入后续case处理
    default:
        println("c")
    }

    b := 0
    switch {                //switch不指定条件表达式或直接为true，可用于替代if ... else if ... else ..
    case b > 1: println("d")
    case b > 2: println("e")
    default: println("f")
    }
}

/*
    break continue goto
*/
func other_control() {

L1:
    for {
        for i := 0; i < 10; i++ {
            if i > 2 {
                break L1
            }else if i < 2{
                println("L1:", i)
            }else {
                goto end            //建议往后goto
            }
        }
    }

L2:
    for i := 0; i < 5; i++ {
        for {
            println("L2:", i)
            continue L2
        }
    }
end:
    println("over")
}

func main() {
    if_control()
    for_control()
    switch_control()
    other_control()
}
