package main

import (
    "fmt"
)

func main() {
    var num, succ, coppie int
    fmt.Scan(&num)
    for num!=0{
      fmt.Scan(&succ)
      if succ==0{
        break
      }
      if succ-num==1{
        coppie++
      }
      num=succ
    }
    fmt.Println(coppie)
}
