package main

import (
    "fmt"
)

//questa funzione mi restituisce il valore presente al n-esimo posto nella serie di Fibonacci
func f_rec (n int) uint64 {
 if n == 1 || n == 2 {
 return 1
}
 return f_rec (n -1) + f_rec (n -2)
}


func f_iter1 (n int) uint64 {
 var f uint64
 var f1 , f2 uint64 = 1, 1

 if n == 2 || n == 1 {
 return 1
 }

 for n >= 3 {
 n--
 f = f1 + f2
 f1 = f2
 f2 = f
}
  return f
}

func f_iter2 (n int) uint64 {
 var f uint64
 var f1 , f2 uint64 = 1, 1

 if n == 2 || n == 1 {
 return 1
 }

 for i := 2; i < n; i ++ {
 f = f1 + f2
 f1 = f2
 f2 = f
}
 return f
}

func f_riter (a , b uint64, n int) uint64 {
 if n == 2 {
 return a
}

 if n == 1 {
 return b
}

 return f_riter (a+b , a , n -1)
}




func main() {
  var n int
  fmt.Scan(&n)
  fmt.Println(f_rec(n))
  fmt.Println(f_riter(1,1,n))
}
