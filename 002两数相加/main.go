package main

import (
	"fmt"
	"log"
)

// 一个节点，除了自身的数据之外，还必须指向下一个节点，尾部节点指向为nil
type LinkNode struct {
	Payload int    //  Payload 为任意数据类型
	Next *LinkNode
}

type LinkNoder interface {
	// go语言接口，在这个接口里面，我们可以定义一系列的方法。
	Add(payload int)
	Delete(index int) int
	Insert(index int, payload int)
	GetLength() int
	Search(payload int) int
	GetAll(index int) int
	Reverse() *LinkNode
}

// go语言方法，对比与下面的NewLinkNode，方法可以理解为面向对象里面对象的方法，虽然实现的功能
// 跟函数类似，但是方式是绑定在对象上的，也就是说我们此处的Add是绑定与head这个LinkNode对象的。
// 这个是go语言区别与其他语言的设计方式，也是go语言很重要的一部分。
func (head *LinkNode) Add(payload int)  {
	// 这里采用尾部插入的方式，给链表添加元素
	point := head

	for point.Next != nil{
		point = point.Next
	}
	newNode := LinkNode{payload, nil}
	point.Next = &newNode

	// 头部插入
	//newNode := LinkNode{payload, nil}
	//newNode.Next = head
}

// 删除元素，并且返回删除的节点的值
func (head *LinkNode) Delete(index int) int  {
	// 边界条件
	linkLength := head.GetLength()

	if index < 0 || index > linkLength{
		fmt.Printf("index out of range %d, please check.", linkLength)
		return -1
	}

	point := head
	for i := 0; i<index; i++{
		point = point.Next     // 移动point到index位置，然后将其next指向next.next
	}

	point.Next = point.Next.Next

	data := point.Next.Payload
	return data
}

// 插入元素
func (head *LinkNode) Insert(index int, payload int){
	linkLength := head.GetLength()

	if index < 0 || index > linkLength{
		fmt.Printf("index out of range %d, please", linkLength)
		return
	}

	point := head
	for i:=0; i<index; i++{
		point = point.Next
	}

	newNode := LinkNode{Payload:payload}
	newNode.Next = point.Next
	point.Next = &newNode
}

func (head *LinkNode) Search(payload int) int {
	point := head

	index := 0

	for point.Next != nil{
		if point.Payload == payload{
			return index
		}else {
			index ++
			point = point.Next

			// 边界条件
			if index > head.GetLength() -1 {
				break
			}

			continue
		}
	}

	// 判断最后一个元素是否匹配
	if point.Payload == payload {
		return index
	}

	return -1   // 不存在时的返回值
}

func (head *LinkNode) GetLength() int  {
	iterator := head

	var length int
	for iterator.Next != nil{
		length ++
		iterator = iterator.Next
	}
	return length
}

func (head *LinkNode) GetAll() []int  {
	dataList := make([]int, 0, head.GetLength())

	point := head
	for point.Next != nil{
		dataList = append(dataList, point.Payload)
		point = point.Next
	}

	dataList = append(dataList, point.Payload)
	return dataList
}

// 创建一个新的链表。 函数，对比与上面的方法，函数是没有绑定任何对象的。
// go语言的函数需要指明参数跟返回值，在此函数中，我们的参数是length，返回值是一个LinkNode对象
// 除了绑定之外，函数跟方法并没有什么不同
func NewLinkNode(length int) *LinkNode  {
	if length <= 0{
		fmt.Printf("链表长度必须大于0")
		log.Panic("链表长度必须大于0")
	}
	var head *LinkNode

	head = &LinkNode{}

	for i := 0; i<length; i++{
		var newNode *LinkNode
		newNode = &LinkNode{Payload: i}
		newNode.Next = head
		head = newNode
	}
	return head
}

// 循环的方式反转一个链表
func (head *LinkNode) Reverse() *LinkNode  {
	if head == nil || head.Next == nil{
		return head
	}

	var reverseHead *LinkNode
	var p *LinkNode

	reverseHead = head
	head = head.Next

	reverseHead.Next = nil

	p = head.Next

	for head != nil{
		head.Next = reverseHead
		reverseHead = head

		head = p
		if p != nil{
			p = p.Next
		}
	}

	return reverseHead
}

// 递归的方式实现链表的反转
func (head *LinkNode) RecursiveReverse() *LinkNode  {
	if head == nil || head.Next == nil{
		return head
	}

	second := head.Next
	newHead := second.RecursiveReverse()

	second.Next = head
	head.Next = nil
	return newHead
}

//实现两个链表数字取和
func addTowNum(t1 *LinkNode, t2 *LinkNode) *LinkNode {
	if t1 == nil && t2 == nil{
		return nil
	}
	if t1 == nil{
		return t2
	}
	if t2 == nil{
		return t1
	}
	
	sum := t1.Payload + t2.Payload
	nextNode := addTowNum(t1.Next, t2.Next)
	if sum < 10{
		return &LinkNode{Payload: sum, Next: nextNode}
	}else {
		tempNode := LinkNode{Payload: 1, Next: nil}
		return &LinkNode{Payload: sum - 10,Next: addTowNum(nextNode, &tempNode)}
	}
	
}

func main()  {
	t1 := LinkNode{Payload:2,Next:nil}
	t1.Add(4)
	t1.Add(3)
	t2 := LinkNode{Payload: 5}
	t2.Add(6)
	t2.Add(4)
	//fmt.Println(t.GetLength())
	fmt.Println(t1.GetAll())
	fmt.Println(t2.GetAll())
	t3 := addTowNum(&t1, &t2)
	fmt.Println(t3.GetAll())
}