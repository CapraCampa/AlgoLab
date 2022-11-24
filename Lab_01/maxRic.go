package main

import (
    "fmt"
    "math"
)

func largest ( numbers []int) int {
  n := len( numbers )
  if n==1 {
    return numbers [0]
  }
  return int(math.Max( float64(numbers [n -1]) , float64(largest ( numbers[:n-1] ))))
}

func main() {
  var n int
  var numbers []int
  for{
    _,err:=fmt.Scan(&n)
    if err!=nil{
      break
    }
    numbers=append(numbers,n)
  }
  fmt.Println(largest(numbers))
}
