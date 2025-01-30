package main

import (
    "fmt"
    "reflect"
)

// demo1 ： 0.3 精度丢失
func demo1() {
    x1 := float64(0.4)
    y1 := float64(0.5)
    fmt.Println("x1+y1=", x1+y1) // x1+y1= 0.9
    x2 := float64(0.3)
    y2 := float64(0.6)
    fmt.Println("x2+y2=", x2+y2) // x+y= 0.899999999999999
}

// demo2 ：常量编译时估值，不会精度丢失，变量会
func demo2() {
    x3 := float64(0.1)
    y3 := float64(0.2)
    fmt.Println("x3+y3=", x3+y3, "0.1+0.2=", 0.1+0.2) // x3+y3= 0.30000000000000004 0.1+0.2= 0.3

    // 如果使用常量，这种情况就不会发生精度损失。
    const c1 = 0.3
    const c2 = 0.6
    fmt.Println("c1+c2=", c1+c2, "c1 type:", reflect.TypeOf(c1))
}

// demo3 ：浮点数不满足结合率
func demo3() {
    var x = 0.3
    var y = 0.6
    var z float64 = 0.1
    fmt.Println("(x+y)+z:", (x+y)+z, "x+(y+z):", x+(y+z)) // (x+y)+z: 0.9999999999999999 x+(y+z): 1
}

// demo4 ：四舍但不一定五入
func demo4() {
    fmt.Printf("%.2f\n", 9.824)  // 9.82
    fmt.Printf("%.2f\n", 9.8250) // 9.82
    fmt.Printf("%.2f\n", 9.8251) // 9.83
}

func main() {
    // fmt.Println(strconv.IntSize)
    // fmt.Println(math.MaxInt16) //32767
    // fmt.Println(math.MinInt16) //-32768

    demo4()

    // // 注意精度问题 ------------------------------------
    // fmt.Printf("%f\n", float32(1)/float32(3)) // 0.333333
    // fmt.Printf("%f\n", float32(1/3))          // 0.000000
    //
    // var f1 float32 = -1.123456789
    // f2 := -1.123456789
    // fmt.Println("f1:", f1, "f2:", f2) // f1: -1.1234568 f2: -1.123456789
}
