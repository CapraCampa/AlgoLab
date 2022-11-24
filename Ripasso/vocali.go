package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func prima(riga string)string{
  for _,char := range riga {
    if strings.ContainsAny(string(char), "aeiouAEIOU"){
      return string(char)
    }
  }
  return "la parola non contiene vocali"
}

func main() {
  scanner:=bufio.NewScanner(os.Stdin)
  for scanner.Scan(){
    fmt.Println(prima(scanner.Text()))
  }
}
