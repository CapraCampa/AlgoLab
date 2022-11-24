package main

import (
    "fmt"
)

func sassi(height int) int{
  if height==1{
    return 1
  }
  return (height*height + sassi(height-1))
}

func main() {
  var n int
  fmt.Scan(&n)
  fmt.Println(sassi(n))
}
