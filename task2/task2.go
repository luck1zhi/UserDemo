package main

import "fmt"

type ListNode struct {
	Next *ListNode
	Val int
}

func NewListNode(val int) *ListNode {
	return &ListNode{Next: nil, Val: val}
}

var HEAD *ListNode
var POINT *ListNode
var END *ListNode

func getLinkedList(){
	fmt.Println("链表值数量：")
	var count int = 0
	fmt.Scanln(&count)


	pos := -1
	fmt.Println("请输入尾部链接的索引位置：")
	fmt.Scan(&pos)

	if pos < -1 || pos > count - 1{
		//输入有误返回-2
		panic("构造有误！")
	}

	fmt.Println("输入链表值构造链表：")
	//定义一个local标记当前链表位置
	var local *ListNode

	for i:=0;i<count;i++{
		val := 0
		fmt.Scan(&val)
		if i == 0{
			//如果是第一个节点的话，记录进HEAD和local
			HEAD = NewListNode(val)
			local = HEAD
			if i == pos{
				POINT = local
			}
		}else{
			local.Next = NewListNode(val)
			local = local.Next
			if i == pos {
				POINT = local
			}
		}
		if i == count - 1{
			END = local
		}
	}
	//检测构造情况
	first := HEAD
	fmt.Println("没有进行链接的构造情况：")
	for ;first!=nil;first=first.Next{
		fmt.Print(first.Val,"  ")
	}
	fmt.Println()

	if pos == -1{
		//没有环直接返回完成构造
		return
	}

	END.Next = POINT
}

func getPos(head *ListNode) *ListNode {

	var pos1 *ListNode = head
	var pos2 *ListNode = head
	var result *ListNode = head

	//确保可以访问Next
	if pos1 == nil || pos1.Next == nil {
		return nil
	}

	for ;pos1!=nil&&pos2.Next!=nil; {
		pos1 = pos1.Next
		pos2 = pos2.Next.Next
		if pos1 == pos2{
			//第一次相遇,result出发
			for ;result != pos1; {
				pos1 = pos1.Next
				result = result.Next
			}
			return result
		}
	}
	return pos2.Next
}

func main() {
	getLinkedList()

	var result *ListNode = getPos(HEAD)
	
	if result == nil{
		//返回空代表没有环
		fmt.Println("该链表没有环！")
		return
	}

	fmt.Println("成环节点：",result)
	fmt.Println("设置节点：",POINT)
}

/**
输出样例：

链表值数量：
5
请输入尾部链接的索引位置：
4
输入链表值构造链表：
1 2 3 4 5
没有进行链接的构造情况：
1  2  3  4  5
成环节点： &{0xc000010250 5}
设置节点： &{0xc000010250 5}

 */

