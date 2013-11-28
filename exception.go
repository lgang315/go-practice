package main

import (
    "fmt"
    "log"
)

//panic触发异常传递，recover仅在defer函数中使用才会终止错误
//todo: 程序执行将终止，如何做到catch的效果？
func pan_rec() {
    defer func() {
        if err := recover(); err != nil {
            log.Fatalln(err)
        }
    }()

    panic("abc")
}

func defer_error() {
    a, b := 10, 0
    defer func() {
        log.Fatalln(recover())
    }()

    defer func() {
        fmt.Println(a / b)  //产生错误，覆盖之前的
    }()

    panic("error")          //被后续错误覆盖
}

func main() {

    pan_rec()
    defer_error()

    //if err := recover(); err != nil {       //不在defer中，错误继续传递
    //    log.Fatalln(err)
    //}

    println("end")
}
