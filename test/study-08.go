package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[int]string{}
	m[1] = "ok"
	m[2] = "no"
	fmt.Println(m)
	fmt.Println(m[1])
	delete(m, 2)
	fmt.Println(m[2])

	mm := map[int]map[int]string{}
	mm[1] = map[int]string{}

	a, ok := mm[2][1]
	//map一定要做检查和MAKE 否则会发生运行时错误
	if !ok {
		mm[2] = map[int]string{}
	}

	mm[2][1] = "good ok"
	a, ok = mm[2][1]

	fmt.Println(a, ok)

	sm := make([]map[int]string, 5)
	for i := range sm {
		sm[i] = make(map[int]string, 1)
		sm[i][1]="ok"
		//smVal 只会做值拷贝 而不是引用
		fmt.Println(sm[i])
	}
	fmt.Println(sm)


	mmap := map[int]string{1:"a", 2:"b",3:"c",4:"d"}//map 无序，但是可以通过数组间接排序
	ss := make([]int, len(mmap))
	i := 0
	for kkk,_ := range mmap  {
		ss[i] = kkk
		i++
	}
	sort.Ints(ss)
	fmt.Println(ss)


	//map sort


	//------homework--//

	homeM1 := map[int]string{1:"a",2:"b",3:"c"}
	//key 和VAL对掉
	//homeM2 := map[string]int{"a":1, "b":2, "c":3}
	homeM3 := make(map[string]int, len(homeM1))

	homeI := 1
	for homeK,_ := range homeM1  {
		homeM3[homeM1[homeK]] = homeK
		homeI++
	}

	fmt.Println(homeM1)
	fmt.Println(homeM3)

}
