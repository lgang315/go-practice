package main


//接收多个参数，无返回值
func test(a, b int , c string) {
    println(a, b, c)
}

//单个返回值
func add(a, b int) int {
    return a + b
}

type callback func(a, b int) int    //定义函数类型

func callfunc(a, b int, cb callback) int{
    return cb(a, b)
}

//命名返回参数
func change(a, b int) (x, y int) {
    x = a + 1
    y = b + 1
    //return y, x                   //按return顺序返回
    return                          //return 空，则按命名参数顺序返回
}

//变参
func sum(s string, args ...int) int {
    var x int
    for _, n := range args {
        x += n
    }
    println(s, x)
    return x
}

//闭包支持
func closure(x int) (func(int) int) {
    //返回匿名函数
    return func (y int) int {
        return x + y
    }
}

//defer，函数退出清场函数，多个defer按FILO执行
//常用与资源清理、关闭文件、解锁、记录执行时间
func f_defer(a, b int) (d int) {
    defer println("defer1:", a, "/", b)
    defer func() {println("return :", d)}()   //调用在ret之前执行，所以如果return修改命名返回值，则闭包中引用的结果就是修改后的值
    defer func() {
        println("defer2:", a, b)
    }()                         //匿名函数，参数列表为空

    return a / b                //执行时若出现严重error，也会被执行
}

//闭包 defer
func f_closure_defer() {
    var fs = [4]func(){}
    for i := 0; i < 4; i++{
        defer println("defer i =", i)                       //值拷贝，将直接获取当前i值，即0，1，2，3
        defer func(){ println("defer_closure ", i)}()       //闭包引用对象，执行时才获取，都是4
        fs[i] = func(){println("closure i = ", i)}          //持有i的引用，执行再获取，都是4
    }
    for _, f := range fs { f() }
}

func main() {
    test(1,2,"adf")
    println(add(1, 2))
    println(add(change(3, 4)))      //多参数返回 作为函数参数

    //var f callback = add
    f := add
    println(callfunc(1, 4, f))

    sum("1 + 2 + 3 = ", 1, 2, 3)
    x := []int{1, 2, 3}
    sum("1 + 2 = ", x[:2]...)       //注意用...展开传参

    fc := closure(10)               //匿名函数，引用了closure.x，每次调用都会产生一个新的匿名函数对象
    println(fc(1))                  //11
    println(fc(2))                  //12

    f_defer(10, 2)
    //f_defer(10, 0)                  //出现严重error，defer也会执行

    f_closure_defer()
}




