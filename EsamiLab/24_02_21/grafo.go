package main

import("fmt")

type node struct{
	next *node
	val int
}

type linkedList struct{
	head *node
}

func newNode(n int)*node{
	return &node{nil,n}
}

func addNewNode(l *linkedList, i int){
	n:=newNode(i)
	if l.head==nil{
		l.head=n
	}else{
		p:=l.head
		for p.next!=nil{
			p=p.next
		}
		p.next=n
	}
}

func stampa(l linkedList){
	if l.head==nil{
		return
	}
	p:=l.head
	for p!=nil{
		fmt.Print(p.val," ")
		p=p.next
	}
}

func main(){
	var l linkedList
	addNewNode(&l,2)
	addNewNode(&l,3)
	stampa(l)	
}
