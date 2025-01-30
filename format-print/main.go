package main

import (
    "bufio"
    "fmt"
    "os"
)

var (
    appName string
    version string
)

func main() {
    // 布尔类型
    // ok := true
    // fmt.Printf("%T,%t \n", ok, ok)

    // var r rune = 65
    // //整数类型
    // // %d 表示以十进制格式输出整数
    // fmt.Printf("%T, %d \n", 123456789, 123456789)   //int, 123456789
    //
    // // %5d 表示以十进制格式输出整数，并且至少占用5个字符的宽度，不足部分用空格填充
    // fmt.Printf("%T, %5d \n", 123456789, 123456789)  //int, 123456789
    //
    // // %5d 同上，但这里的整数12不足5位，所以前面会用空格填充
    // fmt.Printf("%T, %5d \n", 12, 12)                //int,    12
    //
    // // %05d 表示以十进制格式输出整数，并且至少占用5个字符的宽度，不足部分用0填充
    // fmt.Printf("%T, %05d \n", 123456789, 123456789) //int, 123456789
    //
    // // %05d 同上，但这里的整数12不足5位，所以前面会用0填充
    // fmt.Printf("%T, %05d \n", 12, 12)               //int, 00012
    //
    // // %b 表示以二进制格式输出整数
    // fmt.Printf("%T, %b \n", 123456789, 123456789)   //int, 111010110111100110100010101
    //
    // // %o 表示以八进制格式输出整数
    // fmt.Printf("%T, %o \n", 123456789, 123456789)   //int, 726746425
    //
    // // %c 表示输出整数对应的Unicode字符
    // fmt.Printf("%T, %c \n", 66, 66)                 //int, B
    //
    // // %q 表示输出整数对应的Unicode字符，并用单引号括起来
    // fmt.Printf("%T, %q \n", 66, 66)                 //int, 'B'
    //
    // // %x 表示以十六进制格式输出整数，字母小写
    // fmt.Printf("%T, %x \n", 123456789, 123456789)   //int, 75bcd15
    //
    // // %X 表示以十六进制格式输出整数，字母大写
    // fmt.Printf("%T, %X \n", 123456789, 123456789)   //int, 75BCD15
    //
    // // %U 表示输出Unicode码点，格式为 U+XXXX
    // fmt.Printf("%T, %U \n", '中', '中')               //int32, U+4E2D //字符的字面量是rune类型
    //
    // // %v 表示输出变量的默认格式，%s 表示输出字符串
    // fmt.Printf("%T, %v ,%s \n", r, r, string(r))    //int32, 65 ,A
    //
    // // %c 表示输出整数对应的Unicode字符
    // fmt.Printf("%T, %c \n", r, r)                   //int32 ,A

    // var s = "面试"
    // fmt.Println(s[0])
    // for i, item := range s {
    //	//0 : A
    //	//1 : B
    //	fmt.Println(i, ":", item)
    //	fmt.Printf("i:%c\n", item)
    // }

    // 浮点型
    // // %b 表示以科学计数法（二进制）的形式输出浮点数，格式为：尾数p指数
    // fmt.Printf("%b \n", 1000.123456789)   //8797178959608267p-43
    //
    // // %f 表示以十进制浮点数格式输出，默认保留6位小数
    // fmt.Printf("%f \n", 1000.123456789)   //1000.123457
    //
    // // %f 同上，但这里的浮点数1000.0没有小数部分，所以补全6位小数
    // fmt.Printf("%f\n", 1000.0)            //1000.000000
    //
    // // %.2f 表示以十进制浮点数格式输出，保留2位小数
    // fmt.Printf("%.2f \n", 1000.123456789) //1000.12
    //
    // // %.2f 同上，但这里的浮点数1000.125会四舍五入到2位小数
    // fmt.Printf("%.2f \n", 1000.125)       //1000.12
    //
    // // %.2f 同上，但这里的浮点数1000.126会四舍五入到2位小数
    // fmt.Printf("%.2f \n", 1000.126)       //1000.13
    //
    // // %e 表示以科学计数法（小写e）的形式输出浮点数
    // fmt.Printf("%e\n", 1000.1234567898)   //1.000123e+03
    //
    // // %.5e 表示以科学计数法（小写e）的形式输出浮点数，保留5位小数
    // fmt.Printf("%.5e\n", 1000.1234567898) //1.00012e+03
    //
    // // %E 表示以科学计数法（大写E）的形式输出浮点数
    // fmt.Printf("%E\n", 1000.1234567898)   //1E+03
    //
    // // %.5E 表示以科学计数法（大写E）的形式输出浮点数，保留5位小数
    // fmt.Printf("%.5E\n", 1000.1234567898) //1.00012E+03
    //
    // // %F 表示以十进制浮点数格式输出，与 %f 相同
    // fmt.Printf("%F \n", 1000.123456789)   //1000.123457
    //
    // // %g 表示根据数值的大小自动选择 %f 或 %e 格式输出，以更紧凑的形式显示
    // fmt.Printf("%g \n", 1000.123456789)   //1000.123456789
    //
    // // %G 表示根据数值的大小自动选择 %F 或 %E 格式输出，以更紧凑的形式显示
    // fmt.Printf("%G \n", 1000.123456789)   //1000.123456789

    // 字符串
    // arr := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
    // arr1 := []byte{103, 111, 108, 97, 110, 103}
    // arr2 := []uint8{103, 111, 108, 97, 110, 103}
    // arr3 := []rune{'g', 'o', 'l', 'a', 'n', 'g'}
    // arr4 := []int32{'g', 'o', 'l', 'a', 'n', 'g'}
    //
    // // %s 表示输出字符串的内容
    // fmt.Printf("%s \n", "go面试专题")       //go面试专题
    //
    // // %q 表示输出字符串的内容，并用双引号括起来
    // fmt.Printf("%q \n", "go面试专题")       //"go面试专题"
    //
    // // %x 表示将字符串的每个字节转换为小写的十六进制格式输出
    // fmt.Printf("%x \n", "go面试专题")       //676fe99da2e8af95e4b893e9a298
    //
    // // %X 表示将字符串的每个字节转换为大写的十六进制格式输出
    // fmt.Printf("%X \n", "go面试专题")       //676FE99DA2E8AF95E4B893E9A298
    //
    // // %s 表示输出字节数组的内容（字节数组会被当作字符串处理）
    // fmt.Printf("%T, %s \n", arr, arr)   //[]uint8, golang
    //
    // // %q 表示输出字节数组的内容，并用双引号括起来（字节数组会被当作字符串处理）
    // fmt.Printf("%T, %q \n", arr, arr)   //[]uint8, "golang"
    //
    // // %x 表示将字节数组的每个字节转换为小写的十六进制格式输出
    // fmt.Printf("%T, %x \n", arr, arr)   //[]uint8, 676f6c616e67
    //
    // // %X 表示将字节数组的每个字节转换为大写的十六进制格式输出
    // fmt.Printf("%T, %X \n", arr, arr)   //[]uint8, 676F6C616E67
    //
    // // %s 表示输出字节数组的内容（字节数组会被当作字符串处理）
    // fmt.Printf("%T, %s \n", arr1, arr1) //[]uint8, golang
    //
    // // %s 表示输出字节数组的内容（字节数组会被当作字符串处理）
    // fmt.Printf("%T, %s \n", arr2, arr2) //[]uint8, golang
    //
    // // %c 表示将整数数组的每个元素转换为对应的Unicode字符输出
    // fmt.Printf("%T, %c \n", arr3, arr3) //[]int32, [g o l a n g]
    //
    // // %c 表示将整数数组的每个元素转换为对应的Unicode字符输出
    // fmt.Printf("%T, %c \n", arr4, arr4) //[]int32, [g o l a n g]

    // 指针
    // var name *string
    // tmp := "go面试"
    // name = &tmp
    // %p 表示输出指针的地址（以十六进制格式表示）
    // fmt.Printf("%p \n", name) //0xc00009e220

    // Scan 遇到遇到空白字符（空格、制表符、换行符）结束
    // var (
    //     appName string
    //     version int
    // )
    // fmt.Println("请输入name:")
    // fmt.Scan(&appName)
    // fmt.Println("请输入version")
    // fmt.Scan(&version)
    // fmt.Printf("fmt.Scan appName:%s version:%d \n", appName, version)

    // Scanf
    // _, err := fmt.Scanf("name=%s ver=%s", &appName, &version)
    // if err != nil {
    //     fmt.Println(err)
    // }
    // fmt.Printf("fmt.Scanf appName:%s version:%s \n", appName, version)

    // Scanln 遇到换行符（\n）结束
    // fmt.Println("请输入name")
    // fmt.Scanln(&appName)
    // fmt.Println("请输入version")
    // fmt.Scanln(&version)
    // fmt.Printf("fmt.Scan appName:%s version:%s \n", appName, version)

    reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
    fmt.Println("请输入：")
    text, _ := reader.ReadString('\n') // 读到换行
    // text, _ := reader.ReadString(' ') // 读到换行
    fmt.Println(text)

    // go run -ldflags "-X 'main.appName=test' -X 'main.version=1'" main_test.go
    // // go build -ldflags "-X 'main.appName=test' -X 'main.version=1'" main_test.go
    // fmt.Printf("appName:%s version:%s \n", appName, version)
}
