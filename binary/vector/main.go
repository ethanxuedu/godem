package main

import "fmt"

/*
按位移动：
向右移动一位则除以 2，移动两位则除以 2 的平方
向左移动一位则乘以 2，移动两位则乘以 2 的平方
*/

/**
 * @Description: 向左移位
 * @param num- 等待移位的十进制数, m- 向左移的位数
 * @return int- 移位后的十进制数
 */
func leftShift(num int, m int) int {
    return num << m
}

/**
 * @Description: 向右移位
 * @param num- 等待移位的十进制数, m- 向右移的位数
 * @return int- 移位后的十进制数
 */
func rightShift(num int, m int) int {
    return num >> m
}

func main() {
    num := 16
    leftShiftResult := leftShift(num, 2)
    rightShiftResult := rightShift(num, 2)

    fmt.Printf("Original number: %d\n", num)
    fmt.Printf("After left shift by 2: %d\n", leftShiftResult)
    fmt.Printf("After right shift by 2: %d\n", rightShiftResult)
}
