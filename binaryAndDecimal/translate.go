package tanslate

import "fmt"

// 二进制转十进制
func binaryToDecimal(p int) int {
	return 0
}

// 反转字符串
func reverseString(s string) string {
    runes := []rune(s)
    for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
        runes[from], runes[to] = runes[to], runes[from]
    }
    return string(runes)
} 

// 十进制转二进制
func decimalToBinary(p int) string {
	
	var remainder int;
	var binary string;

	for true {
		p = p / 2;
		remainder = p % 2

		binary = binary + fmt.Sprintf("%d", remainder)
		
		if p == 1 || p == 0 {
			binary = binary + fmt.Sprintf("%d", p)
			break
		}
	}

	return binary
}