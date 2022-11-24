package main

import (
    "fmt"
)

func main() {
    var quanti int
    fmt.Scan(&quanti)
    giorni:=make([]int,quanti)
    for i := 0;i < quanti ;i++ {
      fmt.Scan(&giorni[i])
    }
    for i := 1;i < quanti-1;i++ {
      if giorni[i]<giorni[i-1] && giorni[i]<giorni[i+1]{
        fmt.Printf("Ha piovuto il %d giorno\n",i+1)
      }
    }
}
