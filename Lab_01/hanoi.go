package main

import (
    "fmt"
)

func hanoi( n , from, temp, to int ){

  if n!=1{
    hanoi(n-1,from,to, temp)
    fmt.Printf("%d -> %d\n",from,to)
    hanoi(n-1, temp, from, to)
  }else{
    fmt.Printf("%d -> %d\n",from,to)
  }

}

func main() {
  var n int
  fmt.Scan(&n)
  hanoi(n,0,1,2)
}
