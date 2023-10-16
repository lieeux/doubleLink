package tools

import "fmt"

type Node struct {
	Id   int
	Name string
	pre  *Node //指向前一个节点
	next *Node //指向后一个节点
}

func InsertNode(head *Node, newNode *Node) { //在链表末尾添加节点
	temp := head //创建一个辅助节点
	for {        //找链表尾
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = newNode //将链表末尾向后指向新节点
	newNode.pre = temp  //将新节点向前指向链表末尾
}

func InsertNodeByNum(head *Node, newNode *Node) { //按num顺序插入节点
	temp := head //创建一个辅助节点
	for {
		if temp.next == nil {
			break
		} else if temp.next.Id >= newNode.Id {
			break
		}
		temp = temp.next
	}
	newNode.next = temp.next //先建立新节点的指向
	newNode.pre = temp
	temp.next = newNode      //再把原节点指向新节点，不然链就断了
	if newNode.next != nil { //如果新加的节点在链尾，把它后面的节点向前连接会报错
		newNode.next.pre = newNode
	}
}

func DelNode(head *Node, id int) { //删除节点
	temp := head //创建一个辅助节点
	flag := false
	for {
		if temp.next == nil {
			break
		} else if temp.next.Id == id {
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		temp.next = temp.next.next //把当前节点指向下个节点的下个节点
		if temp.next != nil {      //如果删除的节点在链尾，把它后面的节点向前连接会报错
			temp.next.pre = temp
		}
	} else {
		fmt.Printf("id%d不存在\n", id)
	}

}

func ListNode(head *Node) { //输出链表
	temp := head
	if temp.next == nil {
		fmt.Println("链表为空")
		return
	}
	for {
		fmt.Printf("[%d %s]->", temp.next.Id, temp.next.Name)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
	fmt.Println()
}

func ListNodeReverse(head *Node) { //反向输出链表
	temp := head
	if temp.next == nil {
		fmt.Println("链表为空")
		return
	}
	for {
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
	for {
		fmt.Printf("[%d %s]->", temp.Id, temp.Name)
		temp = temp.pre
		if temp.pre == nil {
			break
		}
	}
	fmt.Println()
}
