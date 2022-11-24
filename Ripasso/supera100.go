package main

import (
    "os"
    "fmt"
    "strconv"
)

func main() {
    numeri:=os.Args[1:]
    for _,num := range numeri {
      n,_:=strconv.Atoi(num)
      if n > 100{
        fmt.Println(n)
        return
      }
    }
    fmt.Println("nessun numero maggiore di 100")
}
