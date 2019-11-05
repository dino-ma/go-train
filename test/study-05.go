package main

import "fmt"

func main() {
	a := 1
	var p *int = &a; //指向变量内存地址
	fmt.Println(p);//输出变量内存地址
	fmt.Println(*p)//输出指针值

	b := 1;
	b++;// ++ --为语句 而不是表达式
	fmt.Println(b);
	if 1<2 {
		fmt.Println("1<2")
	}

	c := 990;
	if c:=3; c>1 {//if else 中的同名变量 内部优先级最高，但是不会覆盖原有值。其他作用域用于原同名变量
		fmt.Println(c);
	}
	fmt.Println(c);


	//for 三种形式
	item :=1
	for {
		item++
		if item > 3{
			break;
		}
		fmt.Println(item)
	}
	fmt.Println("over")
	item2 :=1
	for item2 <=3 {
		fmt.Println(item2)
		item2++
	}
	fmt.Println("2over")

	item3 :=1
	for i:=0; i<3; i++ {
		item3++
		fmt.Println(item3);
	}
	fmt.Println("3over")
}
