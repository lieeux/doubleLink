package main

import (
	"doubleLink/tools"
)

func main() {
	head := &tools.Node{} //头结点

	hero1 := &tools.Node{
		Id:   1,
		Name: "穆",
	}
	hero2 := &tools.Node{
		Id:   2,
		Name: "阿鲁迪巴",
	}
	hero3 := &tools.Node{
		Id:   3,
		Name: "撒加",
	}
	hero4 := &tools.Node{
		Id:   4,
		Name: "加隆",
	}

	tools.InsertNode(head, hero1)
	tools.InsertNode(head, hero2)
	tools.InsertNode(head, hero4)
	tools.InsertNodeByNum(head, hero3)

	tools.ListNode(head)
	tools.ListNodeReverse(head)

	tools.DelNode(head, 5)
}
