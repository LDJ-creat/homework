package main

import (
	// "encoding/json"
	"fmt"
	"strconv"
)

// complex128转string的几种方法
func main() {
	var num complex128 = 3 + 4i

	result := strconv.FormatComplex(num, 'e', -1, 64)
	fmt.Println(result)
	// result, err := json.Marshal(num)
	// if err != nil {
	// fmt.Println(err)
	// }---json无法转换复数---
	s := fmt.Sprintf("%v\n", num)
	fmt.Println(s)

	str := strconv.FormatComplex(num, 'f', 2, 64)
	fmt.Println(str)
	//strconv.FormatFloat(c complex, fmt byte, prec int, bitSize int) string的用法
	/*c表示要转换的复数，
	fmt表示输出的格式；'e'表示科学计数法，‘f’表示十进制浮点数，‘g’表示自动选择格式（视数字大小自动选择f或e），
	prec表示复数的位数，-1表示输出的有效数字位数，2表示保留两位有效数字。f--2表示保留两位小数，e---1表示自动选择小数点位置。
	bitSize表示输出的位数，64表示输出的位数为64位。--位数越高精度越高，所占内存越大*/
}
