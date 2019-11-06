package main

import "fmt"

func main() {

	var a [2]int
	var b [2]int
	b = a
	fmt.Println(a)
	fmt.Println(b)

	array := [2]int{1, 10}
	fmt.Println(array)
	array2 := [20]int{19: 1, 1: 2}
	fmt.Println(array2)

	array3 := [...]int{19: 1, 1: 2}
	array4 := [...]int{12}
	fmt.Println(array3)
	fmt.Println(array4)

	var p *[20]int = &array3
	fmt.Println(p)
	x, y := 1, 2
	z := [...]*int{&x, &y}
	fmt.Println(z)

	//在GO中 数组为值类型，会做内存拷贝。而不是引用。
	t := [2]int{3, 2}
	k := [2]int{1, 2}
	fmt.Println(t == k) //相同类型（相同长度）。数组只可以做值比较 == 或 != 不能用于 < > 大小比较

	f := [10]int{}
	f[1] = 2
	fmt.Println(f)

	s := new([10]int) //用NEW创建会直接返回指向数组的内存指针
	s[1] = 2
	fmt.Println(s)

	//多维数组 2个数组 3个元素
	h := [...][3]int{
		{1, 2, 3},
		{2, 3, 4},
		{2, 3, 4},
		{2, 3, 4},
		{2, 3, 4},
		{2, 3, 4},
	}

	fmt.Println(h)

	//go的冒泡排序
	fmt.Println("冒泡排序开始")

	l := [...]int{1, 2, 3, 234, 234, 23, 42, 34, 23, 4, 23, 434, 534, 5, 34, 5}
	num := len(l)
	fmt.Println(l)
	fmt.Println(num)

	for i := 0; i < num; i++ {
		for j := num - 1; j > i; j-- {
			if l[i] < l[j] {
				temp := l[i]
				l[i] = l[j]
				l[j] = temp
			}
		}
	}

	fmt.Println(l)

}
