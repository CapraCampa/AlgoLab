package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)

type listNode struct {
  item int
  next *listNode
}

type linkedList struct {
  head *listNode
}

//che crea un nuovo nodo di lista
func newNode(n int) *listNode{
  node := new( listNode )
  node.item = n
  return node
}

//, che inserisce un nuovo nodo in testa alla lista;
func addNewNode(l linkedList, n int)linkedList{
  node:=newNode(n)
  node.next=l.head
  l.head=node
  return l
}

//, che stampa una lista;
func printList(l linkedList){
  p:=l.head
  for p != nil{
    fmt.Print(p.item," ")
    p=p.next
  }
  fmt.Println()
}

//, che cerca un elemento in una lista;
func searchList(l linkedList, n int)(bool,*listNode){
  p:=l.head
  for p!=nil{
    if p.item==n{
      return true,p
    }
    p=p.next
  }
  return false,nil
}

//, che cancella un item da una lista.
func removeItem(l linkedList, n int) linkedList{
  var prev *listNode=nil
  p:=l.head
  for p.item!=n{
    prev=p
    p=p.next
  }
  if prev==nil{
    l.head=p.next
  }else{
    prev.next=p.next
  }
  return l
}

func conta(l linkedList) (c int){
  p:=l.head
  for p!=nil{
    c++
    p=p.next
  }
  return c
}

func main() {
  scanner:=bufio.NewScanner(os.Stdin)
  var l linkedList
  for scanner.Scan(){
    riga:=strings.Fields(scanner.Text())
    switch riga[0]{
    case "+":
      n,_:=strconv.Atoi(riga[1])
      if flag,_:=searchList(l,n); !flag{
        l=addNewNode(l,n)
      }
    case "-":
      n,_:=strconv.Atoi(riga[1])
      if flag,_:=searchList(l,n); flag{
        l=removeItem(l,n)
      }
    case "?":
      n,_:=strconv.Atoi(riga[1])
      fmt.Printf("L'elemento %d ",n)
      if flag,_:=searchList(l,n); flag{
        fmt.Print("non ")
      }
      fmt.Println("fa parte della lista")
    case "c":
      fmt.Println("Lista lunga", conta(l))
    case "p":
      printList(l)
    case "o":
    case "d":
      l.head=nil
      /*
      p:=l.head
      for p!=nil{
        ogg:=p.item
        removeItem(l,ogg)

      }*/
    case "f":
      break
    }
  }
}
