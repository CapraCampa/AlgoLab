package main

import (
    "fmt"
)

func multiply (x , y int) int {
  if y==0 {
    return 0
  } else {
    return x + multiply (y -1 , x)
  }
}

func main() {
  var n1,n2 int
  fmt.Scan(&n1)
  fmt.Scan(&n2)
  fmt.Println(multiply(n1,n2))
}
