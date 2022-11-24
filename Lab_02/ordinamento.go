package main

import (
    "fmt"
)

/*Scrivete una funzione che legga da standard input una sequenza di interi distinti terminati da
0, memorizzandoli in un vettore ordinato (valutate se è più opportuno usare un array o una
slice): ogni volta che viene letto un nuovo intero, il vettore viene scorso fino a trovare l’esatta
collocazione del numero, quindi si crea lo spazio per il nuovo numero spostando in avanti i
numeri successivi già memorizzati.
Questo algoritmo è utile per riempire un vettore mantenendolo ordinato ad ogni passo.*/
func insertionSort() (numeri []int){
  var n int
  for{
    fmt.scan(&n)
    if n==0{
      break
    }
    if len(numeri)==0{
      numeri=append(numeri,n)
    }else{
      if n>numeri[len(numeri)-1]{
        numeri=append(numeri,n)
      }else{
        for pos,num := range numeri {
          if n<num{
            numeri=append(numeri,numeri[len(numeri)-1])
            for i := ;i < count;i++ {

            }
          }
        }
      }
    }
  }

}

func leggi() (numeri []int){
  var n int
  for{
    _,err:=fmt.Scan(&n)
    if err!=nil{
      break
    }
    numeri=append(numeri,n)
  }
  return numeri
}

func main() {
  numeri:=insertionSort()
  fmt.Println(numeri)
}
