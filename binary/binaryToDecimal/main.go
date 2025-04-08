package main

/*
   十进制与二进制互转
*/

import (
    "fmt"
    "strconv"
)

// DecimalToBinary 将十进制数转换为二进制字符串
func DecimalToBinary(decimalSource int) string {
    return strconv.FormatInt(int64(decimalSource), 2)
}

// BinaryToDecimal 将二进制字符串转换为十进制数
func BinaryToDecimal(binarySource string) (int, error) {
    result, err := strconv.ParseInt(binarySource, 2, 64)
    if err != nil {
        return 0, err
    }
    return int(result), nil
}

func main() {
    // 示例：十进制转二进制
    decimalNumber := 42
    binaryString := DecimalToBinary(decimalNumber)
    fmt.Printf("十进制 %d 转换为二进制: %s\n", decimalNumber, binaryString)

    // 示例：二进制转十进制
    binaryString = "101010"
    decimalNumber, err := BinaryToDecimal(binaryString)
    if err != nil {
        fmt.Println("转换失败:", err)
    } else {
        fmt.Printf("二进制 %s 转换为十进制: %d\n", binaryString, decimalNumber)
    }
}
