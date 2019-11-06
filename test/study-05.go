package main

import (
	"fmt"
)

func main() {
	a := 1
	var p *int = &a //指向变量内存地址
	fmt.Println(p)  //输出变量内存地址
	fmt.Println(*p) //输出指针值

	b := 1
	b++ // ++ --为语句 而不是表达式
	fmt.Println(b)
	if 1 < 2 {
		fmt.Println("1<2")
	}

	c := 990
	if c := 3; c > 1 { //if else 中的同名变量 内部优先级最高，但是不会覆盖原有值。其他作用域用于原同名变量
		fmt.Println(c)
	}
	fmt.Println(c)

	//for 三种形式
	item := 1
	for {
		item++
		if item > 3 {
			break
		}
		fmt.Println(item)
	}
	fmt.Println("over")
	item2 := 1
	for item2 <= 3 {
		fmt.Println(item2)
		item2++
	}
	fmt.Println("2over")

	item3 := 1
	for i := 0; i < 3; i++ {
		//i = 3;
		item3++
		fmt.Println(item3)
	}
	fmt.Println("3over")

	fmt.Println("------switch start----")
	//switch
	switchA := 1
	switch switchA {
	case 0:
		fmt.Println("a=0")
	case 1:
		fmt.Println("a=1")
	default:
		fmt.Println("none")
	}

	switch {
	case switchA >= 0:
		fmt.Println("a>=0")
		fallthrough
	case switchA >= 1:
		fmt.Println("a=1")
	default:
		fmt.Println("none")
	}

	switch switchA := 5; {
	case switchA >= 0:
		fmt.Println("a>=0")
		fallthrough
	case switchA >= 1:
		fmt.Println("a=1")
	default:
		fmt.Println("none")
	}

	fmt.Println("--------break---------")
LABLE1:

	for {
		for i := 0; i < 10; i++ {
			if i == 2 {
				goto LABLE2
			}
			if i > 3 {
				break LABLE1
			}
			fmt.Println(i)
		}
	}
LABLE2:
	fmt.Println("for end")

LABLE3:
	for i := 0; i < 10; i++ {
		fmt.Println("for i")
		fmt.Println(i)
		fmt.Println("for i end")
		for {
			continue LABLE3
			fmt.Println(i)
		}
	}
	fmt.Println("lablesok")

	//homework
	//将continue 替换成goto

LABEL:
	for i := 0; i < 10; i++ {
		for {
			fmt.Println(i)
			goto LABEL
		}
	}
	//continue 是继续循环，可以在for循环体中继续循环行为，而goto是将代码重置到标签位置重新进行执行，所以一直输出0 也就是FOR循环中的首次循环

}
