package main

import (
        "fmt"
        "strings"
)

func quanteConA(parole []string) int{
  var numA int
  for _,par := range parole {
    if strings.ContainsRune(par, 'a'){
      numA++
    }
  }
  return numA
}

func primaConA(parole []string) string{
  var prima string
  for _,par := range parole {
    if strings.ContainsRune(par, 'a'){
      prima=par
      break
    }
  }
  return prima
}

func piuCorta(parole []string) int{
  len1:=len(parole[0])
  for _,par := range parole[1:] {
    if len1>len(par){
      len1=len(par)
    }
  }
  return len1
}

func main () {
 const N = 10
 parole := make([]string, N)

 for i := 0; i < N; i ++ {
 fmt . Scan (& parole [i ])
 }

 fmt .Println( quanteConA(parole))
 fmt.Println(piuCorta(parole))
 fmt.Println(primaConA(parole))
}
