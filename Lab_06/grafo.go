package main
import (
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"os"
)

type linkedList struct{
	head *node
}

type node struct{
	val int
	next *node
}

type grafo struct {
	n int // numero di vertici
	adiacenti []*linkedList
}

func nuovoGrafo(n int) (*grafo){
	adiacenti:= make([]*linkedList, n)
	for i := 0; i < n; i++ {
		adiacenti[i]=&linkedList{nil}
	}
	return &grafo{n,adiacenti}
}

func newNode(n int) *node{
	return &node{n, nil}
}

func addNode(n int, l *linkedList){
	if l.head==nil{
		l.head=newNode(n)
		return
	}
	p:=l.head
	for p.next!=nil{
		p=p.next
	}
	p.next=newNode(n)
}

func stampa(l *linkedList){
	p:=l.head
	for p!=nil{
		fmt.Print(p.val, " ")
		p=p.next
	}
}

func trova(v1, v2 int, g grafo)bool{
	return trovaLista(v2, g.adiacenti[v1])
}

func trovaLista(v2 int, l *linkedList)bool{
	p:=l.head
	for p!=nil{
		if p.val==v2{
			return true
		}
		p=p.next
	}
	return false
}
 

func main(){
	scanner:=bufio.NewScanner(os.Stdin)
	var n int
	fmt.Scan(&n)
	fmt.Print("Inserisci archi tra i vertici da 0 a ", n-1, ":\n")
	grafo:=nuovoGrafo(n)
	for scanner.Scan(){
		parole:=strings.Fields(scanner.Text())
		v1,_:=strconv.Atoi(parole[0])
		v2,_:=strconv.Atoi(parole[1])
		addNode(v2,grafo.adiacenti[v1])
	}
	for i,v := range grafo.adiacenti {
		fmt.Print("Il vertice ",i, " ha come vertici di adiacenza: ")
		stampa(v)
		fmt.Println()
	}
	fmt.Println("Inserire due vertici per verificare se sono collegati da un arco:")
	var x,y int
	fmt.Scan(&x)
	fmt.Scan(&y)
	if trova(x,y, *grafo){
		fmt.Printf("I vertici %d e %d sono collegati da un arco\n", x,y)
	}else{
		fmt.Printf("I vertici %d e %d non sono collegati da un arco\n", x,y)
	}
}