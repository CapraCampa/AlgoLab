/*Scrivere un programma che legge da riga di comando un intero che rappresenta il saldo di un conto corrente.
Il programma legge poi da standard input una serie di numeri interi che rappresentano spese da addebitare sul
conto e stampa il saldo finale. La lettura si interrompe quando il saldo è <=0. */
package main

import (
    "fmt"
    "strconv"
    "os"
)

func main() {
  var spesa int
    conto,_:=strconv.Atoi(os.Args[1])
    for conto > 0 {
      fmt.Scan(&spesa)
      conto-=spesa
    }
    fmt.Println(conto)
}
