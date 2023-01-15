package main

import ("fmt")

type Node struct{
	val int
	next *Node
}

type linkedList struct{
	head *Node
}

/*func newNode(v int)*Node{
	return &Node{v,nil}
}

func addNewNode(l *linkedList, val int){
	n:=newNode(val)
	if l==nil{
		l=&linkedList{n}
		return
	}
	p:=l.head
	for p.next!=nil{
		p=p.next
		fmt.Println(p.val)
	}
	p.next=n
}*/

func kLastCount(p *Node, k int)int{
	var count int
	if p==nil{
		return 0
	}
	count=1+kLastCount(p.next,k)
	if count==k{
		fmt.Printf("%d\n",p.val)
	}
	return count
}

func main(){
	var l linkedList
	/*addNewNode(l,3)
	addNewNode(l,2)
	fmt.Println(l.head.val)
	addNewNode(l,5)
	addNewNode(l,1)
	addNewNode(l,7)*/
	l.head=&Node{3,nil}
	l.head.next=&Node{2,nil}
	l.head.next.next=&Node{5,nil}
	l.head.next.next.next=&Node{1,nil}
	l.head.next.next.next.next=&Node{7,nil}
	fmt.Println(kLastCount(l.head,1))
}