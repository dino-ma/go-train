package main

import (
	"fmt"
)

func main() {

	var s1 []int
	fmt.Println(s1)

	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(a)

	s2 := a[5:]
	fmt.Println(s2)


	s3 := make([]int, 10, 100)//类型，填充元素，容量 每次2倍提升
	fmt.Println(s3)
	fmt.Println(len(s3), cap(s3))

	fmt.Println("------------")
	s4 :=[]byte{'a','b','c','d','e','f','g','h','i','j','k','l','m','n'}
	sa := s4[2:5]
	sb := s4[4:6]
	fmt.Println(s4)
	fmt.Println(string(sa))
	fmt.Println(string(sb))


	sc := sa[3:5]
	fmt.Println(string(sc))//切片的容量会到数组的尾部


	s5 := make([]int, 3, 6)
	fmt.Printf("%p \n", s5)

	s5 = append(s1, 1,2,4,5)
	fmt.Printf("%v %p\n", s5, s5)
	s5 = append(s1, 1,2,4,5)
	fmt.Printf("%v %p\n", s5, s5)


	array := []int{1,2,3,4,5,6}
	s6 := array[2:5]
	s7 := array[1:3]
	fmt.Println(s6, s7)
	s7 = append(s7, 1,1,1,1,1,1,1,1,1,1,1,1,1,1,1)//如果切片超过原数组长度则会COPY 内存地址 *2 *2 之后交集部分就不会被覆盖
	s6[0] = 9;
	fmt.Println(s6, s7)



	s8 := []int{1,2,3,4,5,6}
	s9 := []int{7,8,9}
	//copy(s8, s9)
	copy(s9, s8)
	fmt.Println(s8)
	fmt.Println(s9)

	//---homework--//
	s10 := []int{1,2,3,4,5,6,6,7}
	fmt.Println(s10)
	s11 := s10[:]
	fmt.Println(s11)

}

