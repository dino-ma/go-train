package main

import "fmt"

type treeNode struct {
	value int
	left, right * treeNode
}

//接收者
func (node treeNode) print() {
	fmt.Print(node.value)
}

func (node *treeNode) setValue(value int) {
	node.value = value
}

func createNode(value int) *treeNode  {
	return &treeNode{value:value}
}

func main() {




	//var a = createNode(2)
	//fmt.Println(a)
	var root treeNode
	//fmt.Println(root);

	//root = treeNode{value: 3}
	//root.left = &treeNode{}
	//root.right = &treeNode{5, nil, nil}
	//root.right.left = new(treeNode)
	//fmt.Println("left-right")
	//fmt.Println(root);


	//
	//node := []treeNode {
	//	{value: 3},
	//	{},
	//	{6, nil, &root},
	//}
	//fmt.Println(node)
	//
	//fmt.Println(createNode(2))

	root = treeNode{value:3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)

	root.right.left.setValue(4)
	root.right.left.print()
	fmt.Println(root)
}
