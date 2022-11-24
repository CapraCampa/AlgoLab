package main

import (
      "fmt"
        "math"
      )

func stranoProdotto(numeri []int) int{
  prod:=1
  for _,n := range numeri {
    if n>7 && n%4==0{
      prod*=n
    }
  }
  if prod==1{
    return 0
  }
  return prod
}

func pariDispari(numeri []int){
  for _,n := range numeri {
    if n%2==0{
      fmt.Println("pari")
    }else{
      fmt.Println("dispari")
    }
  }
}

func minDispari(numeri []int) int{
  min:=math.MaxInt
  for _,n := range numeri {
    if n%2!=0{
      if min>n{
        min=n
      }
    }
  }
  return min
}

func main () {
 const N = 10
 numeri := make([]int, N)

 for i := 0; i < N; i ++ {
 fmt.Scan (& numeri [i ])
}
  pariDispari(numeri)
  fmt.Println(minDispari(numeri))
  fmt.Println(stranoProdotto(numeri))
}
