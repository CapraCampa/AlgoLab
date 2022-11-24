package main

import (
    "fmt"
    "math/rand"
    "time"
)

type bitreeNode struct {
  left * bitreeNode
  right * bitreeNode
  val int
}

type bitree struct {
  root * bitreeNode
}

func newNode(n int)*bitreeNode{
  node:=new(bitreeNode)
  node.val=n
  return node
}


func arr2tree (a []int, i int) ( root *bitreeNode ) {
  if i >= len(a) {
    return nil
  }
  root=newNode(a[i])
  root.left = arr2tree(a,2*i+1)
  root.right = arr2tree(a,2*i+2)
  return root
}

func stampaAlberoASommario ( node *bitreeNode , spaces int) {
  for i := 0; i < spaces ; i ++ {
    fmt .Print(" ")
  }
  fmt.Print("*")
  fmt.Println(node.val)
  if node . left != nil || node . right != nil {
    stampaAlberoASommario ( node . left , spaces +1)
    stampaAlberoASommario ( node . right , spaces +1)
  }
}

func main() {
  rand.Seed(time.Now().UnixNano())
  len:=(rand.Intn(20))+1
  numeri:=new([]int)
  for i := 0;i < len;i++ {
    numeri=append(numeri,rand.Int())
  }
  root:=arr2tree(numeri,0)
  stampaAlberoASommario(root,0)

}
