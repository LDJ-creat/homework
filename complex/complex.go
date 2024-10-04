package main

import (
	"fmt"
	"hash/fnv"
	"math"
	"strconv"
)

type number complex128

var a, b number
var real1, imag1 float64
var real2, imag2 float64

func input() {
	fmt.Println("请输入第一个复数的实部和虚部，以空格分隔：")
	fmt.Scanln(&real1, &imag1)

	fmt.Println("请输入第二个复数的实部和虚部，以空格分隔：")
	fmt.Scanln(&real2, &imag2)
}

func (n1 number) add(n2 number) {

	n1 = number(complex(real1, imag2))
	n2 = number(complex(real2, imag2))
	sum := complex(real1+real2, imag1+imag2)
	fmt.Println("两复数之和为:", sum)

}

func (n1 number) subtraction(n2 number) {

	n1 = number(complex(real1, imag2))
	n2 = number(complex(real2, imag2))
	difference := number(complex(real1-real2, imag1-imag2))
	fmt.Println("两复数之差为:", difference)

}

func (n1 number) multiplication(n2 number) {

	n1 = number(complex(real1, imag2))
	n2 = number(complex(real2, imag2))
	product := number(complex(real1*real2-imag1*imag2, real1*imag2+real2*imag1))
	fmt.Println("两复数之积为:", product)
}

func (n1 number) division(n2 number) {

	n1 = number(complex(real1, imag2))
	n2 = number(complex(real2, imag2))
	quotient := number(complex(-(real1*real2-imag1*imag2)/(real2*real2+imag2*imag2), ((-real1*imag2)+real2*imag1)/(real2*real2+imag2*imag2)))
	fmt.Println("两复数之商为:", quotient)
}

func (n1 number) magnitude() {
	var real, imag float64
	fmt.Println("请输入复数的实部和虚部，以空格分隔：")
	fmt.Scanln(&real, &imag)
	n1 = number(complex(real, imag))
	magnitude := math.Sqrt(real*real + imag + imag)
	fmt.Println("复数的模为:", magnitude)
}

func (n1 number) toString() string {
	var real, imag float64
	fmt.Println("请输入复数的实部和虚部，以空格分隔：")
	fmt.Scanln(&real, &imag)
	n1 = number(complex(real, imag))
	n1str := strconv.FormatComplex(complex(real, imag), 'f', 0, 64)
	fmt.Println("复数的字符串表示为:", n1str)
	hash := fnv.New32()                           //创建哈希实例
	hash.Write([]byte(n1str))                     //将字符串b写入哈希实例
	hashstring := fmt.Sprintf("%x", hash.Sum32()) //转换为16进制表示的哈希值
	// fmt.Println("复数的哈希值十六进制表示为:", hashstring)
	return hashstring //也可用fmt.Sprintf(%v\n,n1)将复数转换为字符串

}

func menu() {
	fmt.Println("请选择所需的运算：")
	fmt.Println("1.加法")
	fmt.Println("2.减法")
	fmt.Println("3.乘法")
	fmt.Println("4.除法")
	fmt.Println("5.求模长")
	fmt.Println("6.toString(转字符串)")
	fmt.Println("7.退出")
}
func main() {
	var a, b number
	for {
		menu()
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			input()
			a.add(b)
		case 2:
			input()
			a.subtraction(b)
		case 3:
			input()
			a.multiplication(b)
		case 4:
			input()
			a.division(b)
		case 5:
			a.magnitude()
		case 6:
			a.toString()
		case 7:
			fmt.Println("欢迎下次使用")
			return
		default:
			fmt.Println("输入错误，请重新输入！")
		}
	}
}
