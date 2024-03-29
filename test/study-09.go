package main

import (
	"fmt"
)

func main() {

	z, b, c := msjTest(1, "3", 5)
	fmt.Println(z, b, c)

	msjTest2(1, 23, 123, 12, 31, 23, 12, 3, 123, 1, 23, 12, 3, 12, 31, 23, 123, 12, 31, 23, 12, 3, 2) //传到函数中是一个切片。可以按照索引去取

	msjTest3("dino.ma", 1, 123, 12, 3, 12, 3, 12, 3, 1, 23, 1, 2)

	a, d := 1, 2
	msjTest4(a, d)
	fmt.Println(a, d)

	s1 := []int{1, 2, 3, 4}
	msjTest5(s1)
	fmt.Println(s1)

	t := 1
	msjTest6(&t)
	fmt.Println(t)

	u := msjTest8 //函数也是一种变量类型
	u()

	o := func() {
		fmt.Println("我是一个匿名函数")
	}
	o()

	f := closure(10)  //返回函数体
	fmt.Println(f(1)) //调用闭包函数
	fmt.Println(f(2)) //调用闭包函数

	fmt.Println("-----defer---")

	//fmt.Println("a")
	//defer fmt.Println("b")
	//defer fmt.Println("c")
	//for i := 0; i < 3 ; i++  {
	//	defer fmt.Println(i)
	//}
	//fmt.Println("for defer")
	//
	//for i:=0;i<3 ;i++  {
	//	defer func() {
	//		fmt.Println(i)
	//	}()
	//
	//}

	//msjTestA()
	//msjTestB()
	//msjTestC()

	//homework
	var fs = [4]func(){}
	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i=", i) //栈调用  先进后出    先进  后出
		defer func() {
			fmt.Println("defer_closure i =", i)//            后进  先出
		}()
		fs[i] = func() {//first print
			fmt.Println("closure i =", i)
		}
	}

	for _, f := range fs {
		f()

	}
	//这个作业太牛逼了
}

func msjTestA() {
	fmt.Println("Func A")
}

func msjTestB() {
	//如果有PANINC 则recover会返回panic 信息
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover in B panic:", err)
		}
	}()
	panic("Panic in B")

}

func msjTestC() {
	fmt.Println("Func C")
}

func msjTest(a int, b string, c int) (int, string, int) {

	a = 1993
	b = "生日"
	c = 1006

	return a, b, c
}

func msjTest2(a ...int) {
	fmt.Println(a)
}

func msjTest3(str string, param ...int) {
	fmt.Println(str)
	fmt.Println(param)
}

func msjTest4(s ...int) {
	s[0] = 999
	s[1] = 999
	fmt.Println(s)
}

//变量做值拷贝， SLICE切片做内存地址拷贝 所以原SLICE在函数体内如果有改变，则原切片在函数外部也会改变。
func msjTest5(s []int) {
	s[0] = 5
	s[1] = 6
	s[2] = 7
	s[3] = 8
	fmt.Println(s)
}

//引用传递，传递内存地址
func msjTest6(a *int) {
	*a = 7
	fmt.Println(*a)
}

func msjTest8() {
	fmt.Println("Func msjTest8")
}

func closure(x int) func(int) int {
	fmt.Printf("%p\n", &x)
	return func(y int) int {
		fmt.Printf("%p\n", &x)
		return x + y
	}
}
