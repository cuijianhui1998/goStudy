package main

import "fmt"

func intTest() {
	//int分为有符号整型和无符号整型
	//int8 int16 int32 int64 int8的取值范围是-128到127
	//uint8 uint16 uint32 uint64 uint8的取值范围是0-255
	//int uint会根据操作系统的机器字大小在32位和64位间变化
	//uint8也表示byte类型,rune和int32是等价的

	//由于初始化a的时候赋值超过了他的范围,会报错constant 255 overflows int8
	//var a int8
	//a = 255
	//fmt.Print(a)
}

func floatTest() {
	//浮点数分为float32和float64,科学计数法可以用e或E代替,例如1.0e3表示1000.0
}

func complexTest() {
	//复数分为complex64和complex128
}
func boolTest() {
	//布尔值位true和false
}

func strTest() {
	//字符串:双引号包裹起来,可以像python中那样对str进行索引求值,不过返回的是asiic码,
	var s string = "abhsgfs"
	//h的asiic码是104,所以输出的是104
	fmt.Println(s[2])

	//字符串格式化的方式
	//1. 拼接 +
	var m = "acbfh"
	fmt.Print(s + m)
	//2. 使用fmt.Sprint()
	fmt.Print(fmt.Sprint("acbfh%s", s))

	//另外一种定义字符串的方式是使用反引号``,和python中的三个引号的作用有点类似
	//会对字面量中的特殊字符进行转义
	var c = `adc\n\t`
	fmt.Print(c)

	//字符串的方法
	l := len(c) //求长度
	fmt.Print(l)

}

func main() {
	//intTest()
	strTest()
}
