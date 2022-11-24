package main

import (
    "fmt"
    "os"
    "unicode"
    "strings"
)

func maxCon(parole []string)int{
  var max int
  for _,str := range parole {
    quante:=0
    for _,char := range str {
      if unicode.IsLetter(char) && !strings.ContainsAny(string(char),"aeiou"){
        quante++
      }
    }
    if max<quante{
      max=quante
    }
  }
  return max
}

func main() {
  fmt.Println(maxCon(os.Args[1:]))
}
