package main

import (
    "fmt"
)

func main() {
    var prev,succ int
    _,err:=fmt.Scan(&prev)
    somma:=prev
    if err!=nil{
      return
    }
    for{
      _,err:=fmt.Scan(&succ)
      if succ>prev{
        somma+=succ
      }else{
        fmt.Println("La somma è: ",somma)
        somma=succ
      }
      prev=succ
      if err!=nil{
        break
      }
    }
}
