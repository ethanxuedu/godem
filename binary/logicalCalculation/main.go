package main

import "fmt"

/*
逻辑运算
按位“或” (|)：如果两个位中有一个为 1，则该位结果为 1。
按位“与” (&)：如果两个位都为 1，则该位结果为 1。
按位“异或” (^)：如果两个位不同，则该位结果为 1，如果相同，则为 0。
*/

/**
 * @Description: 二进制按位“或”的操作
 * @param num1- 第一个数字，num2- 第二个数字
 * @return 二进制按位“或”的结果
 */
func or(num1 int, num2 int) int {
    return num1 | num2
}

/**
 * @Description: 二进制按位“与”的操作
 * @param num1- 第一个数字，num2- 第二个数字
 * @return 二进制按位“与”的结果
 */
func and(num1 int, num2 int) int {
    return num1 & num2
}

/**
 * @Description: 二进制按位“异或”的操作
 * @param num1- 第一个数字，num2- 第二个数字
 * @return 二进制按位“异或”的结果
 */
func xor(num1 int, num2 int) int {
    return num1 ^ num2
}

func main() {
    num1 := 5 // 0101
    num2 := 3 // 0011

    fmt.Printf("num1: %d, num2: %d\n", num1, num2)
    fmt.Printf("num1 | num2: %d\n", or(num1, num2))  // 0111
    fmt.Printf("num1 & num2: %d\n", and(num1, num2)) // 0001
    fmt.Printf("num1 ^ num2: %d\n", xor(num1, num2)) // 0110
}
