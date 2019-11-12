package main

import "fmt"

type testStruct struct {
}

type human struct {
	Job string
}

type person struct {
	Name string
	Age  int
	Sex  int
}

type student struct {
	StudentName string
	Age         int
	ClassNum    string
}

type man struct {
	Name string
	Age  int
}

type china struct {
	human
	Person string
	Area   string
	Info   struct {
		Name string
		Age  int
	}
}

func main() {
	testTruct := testStruct{}
	fmt.Println(testTruct)

	person := person{}
	person.Age = 19
	person.Name = "msj"
	person.Sex = 1
	fmt.Println(person)

	student := student{
		StudentName: "mashengjie",
		Age:         19,
		ClassNum:    "20199",
	}
	fmt.Println(student)

	testA(student)
	fmt.Println(student)

	testB(&student)
	fmt.Println(student)

	man := &man{
		Name: "马胜杰",
		Age:  123123123,
	}
	fmt.Println(man)
	man.Name = "xxxx"
	fmt.Println(man)
	testC(man)

	x := struct {
		Name string
		Age  int
	}{
		Name: "dino",
		Age:  19,
	}
	fmt.Println(x)

	y := &struct {
		Name string
	}{
		Name: "msj",
	}
	fmt.Println(y)
	y.Name = "yyy"
	fmt.Println(y)

	china := china{
		human:  human{Job: "后端工程师"},
		Area:   "china",
		Person: "100000亿",
	}
	china.Info.Name = "马胜杰"
	china.Info.Age = 19
	fmt.Println(china)
	china.human.Job = "PHP"

	asia := china
	fmt.Println(asia)

}

func testA(stu student) {
	//值拷贝
	stu.Age = 9999
	fmt.Println("stu A:", stu)
}

func testB(stu *student) {
	stu.Age = 88888
	fmt.Println("引用指针:", stu)
}

func testC(man *man) {
	man.Age = 9
	fmt.Println(man)
}
