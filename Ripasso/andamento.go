/*Data da standard input una serie di interi positivi terminata da 0, stampare ’+’ ogni volta che il nuovo valore
è maggiore o uguale al precedente e ’-’ altrimenti.  */
package main

import (
    "fmt"
)

func main() {

    var n, succ int
    fmt.Scan(&n)
    for n!=0 {
      fmt.Scan(&succ)
      if succ==0{
        break
      }
      if succ<n{
        fmt.Println("-")
      }else{
        fmt.Println("+")
      }
      n=succ
    }
}
