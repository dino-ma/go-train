package main

import (
	"fmt"
	"math"
	"os"
)

var name = "马胜杰"

const HOST  = "https://www.baidu.com"

const  (
	A = iota
	B
	C
	D
	E
	F
)

type(
	byte int8
	文本 string//开发中不要用
	byteSize int64//下载
)

func main() {
	age := 18
	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(HOST)
	fmt.Println("hello world")
	fmt.Println(os.Args)

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println("零值：")
	var a int;
	var b bool;
	var c float32;
	var d string;
	var e [1]int;
	var f [1]bool;
	var g [1]byte;
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	//最大值最小值
	fmt.Println(math.MinInt8)
	fmt.Println(math.MaxInt8)

	//开发中不要用
	var bbb 文本
	bbb = "中文类型名牛逼牛逼～"
	fmt.Println(bbb);

	var byteSize byteSize//文件下载大小
	byteSize = 123123;
	fmt.Println(byteSize)


	xxx := false;
	fmt.Println(xxx);


	floatA := 1.1
	intA := int(floatA)
	fmt.Println(intA)

}
