package main

import (
    "fmt"
)

func main() {
    var max int
    var n, nMax string
    for {
      fmt.Scan(&n)
      if n=="-1"{
        break
      }
      zeri:=0
      for _,char := range n {
        if char=='0'{
          zeri++
        }
      }
      if max<=zeri{
        max=zeri
        nMax=n
      }
    }
    fmt.Println(nMax)
}
