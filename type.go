package main
import (
    "fmt"
    "unsafe"
)

func variable() {
    println("//变量声明与初始化")
    var a = 0x1234
    var b string = "hello世界"
    var c, d int = 1, 2     //类型相同
    var e, f = 123, "hello" //自动推断
    g, h := 123, "hello"    //自动推断
    var j bool              //默认值

    var (                   //多个变量类型不同
        p int
        q bool
    )

    println(a, b, c, d, e, f, g, h, j, p, q)

    sa := []int{1, 2, 3}
    i := 0
    i, sa[i] = 1, 2         //先计算所有值，再从左到右赋值 set i = 1, sa[0] = 2
    println(sa[0])

    //var xxx = "error: declared and not used"

    println("//引用类型slice map channel 默认值都是nil")
    var s []int
    var m map[string]int
    var ch chan int
    println(s == nil, m == nil, ch == nil)   //引用类型slice map channel 默认值都是nil

    fmt.Printf("%x\n", uint8(a))            //不支持隐式转换

    println("//常量组，如不提供初始值，则与上行表达式完全相同")
    const (                                 //常量组，如不提供初始值，则与上行表达式完全相同
        c_a = "abc"
        c_b
    )
    println(c_a,c_b)

    println("//枚举值，用iota从0自动增长")
    type ByteSize int64
    const (
        _ = iota                        //忽略
        KB ByteSize = 1 << (10 * iota)  //KB
        MB
        GB
    )
    println(GB)

    println("//`定义原始字符串，不进行转义；+跨行连接必须在上一行行尾；切片返回字串而非slice")
    str := `abcdefg字母\n`
    str1 := str[1:4]
    str_s := "xxx" +
            ", 222"
    println(str, str1, str_s)

    println("//修改字符串，需先转换为[]byte或[]rune,改完还得转回来")
    str_r := []rune(str)
    str_r[1] = '你'         //''表示一个Unicode字符rune
    println(string(str_r), len(str_r), len(str))

    println("//指针，uintptr可以做指针运算；unsafe.Pointer类似void *， 可以在不同指针类型间转换")
    type User struct {
        Id int
        Name string
    }
    x := User{1,"User1"}
    ptr := &x                                       //*User类型的指针
    //ptr++                                         //非法，*int 与 int类型无法运算
    var ptr_u unsafe.Pointer = unsafe.Pointer(ptr)  //*User类型指针转换为unsafe.Pointer
    var np uintptr = uintptr(ptr_u)                 //转换为uintPtr
    np = np + unsafe.Offsetof(ptr.Name)             //指针运算
    var name *string = (*string)(unsafe.Pointer(np))//转换回unsafe.Pointer,再转回*string
    println(*name)

}

func main() {
    variable()
}
