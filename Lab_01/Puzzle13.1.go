package main

import (
    "fmt"
    "strings"
    "strconv"
    "bufio"
    "os"
)

type Punto struct {
  x int
  y int
}

/*func controllo(p Punto, pos int, punti map[Punto]int){
  for pos1,p1 := range punti {
    fmt.Println(p1, p)
    if p1==p && pos!=pos1 {
      if pos1!=(len(punti)-1){
        punti=append(punti[:pos1], punti[pos1+1:]...)
      }else{
        punti=punti[:pos1]
      }
    }
  }
  return punti
}*/

func piega(s string, punti map[Punto]int){
  pieghe:=(strings.Fields(s))[2]
  asse:=strings.Split(pieghe, "=")
  n,_:=strconv.Atoi(asse[1])
  if asse[0]=="x"{
    for punto,_ := range punti {
      if punto.x>n{
        delete(punti, punto)
        punto.x=n-(punto.x-n)
        punti[punto]++
      }
    }
  }else{
    for punto,_ := range punti {
      if punto.y>n{
        delete(punti, punto)
        punto.y=n-(punto.y-n)
        punti[punto]++
      }
    }
  }
}

func main() {
  var punto Punto
  punti:=make(map[Punto]int, 0)
  scanner:=bufio.NewScanner(os.Stdin)
  for scanner.Scan(){
      s:=scanner.Text()
      coordinate:=strings.Split(s,",")
      if s=="z"{
        break
      }
      punto.x,_=strconv.Atoi(coordinate[0])
      punto.y,_=strconv.Atoi(coordinate[1])
      punti[punto]++
    }
    //fmt.Println(punti)
  for scanner.Scan(){
    s:=scanner.Text()
    if s=="z"{
      break
    }
    piega(s,punti)
  }
  var maxX,maxY int
  for p,_ := range punti {
    if p.x>maxX{
      maxX=p.x
    }
    if p.y>maxY{
      maxY=p.y
    }
  }
  for y := 0;y<=maxY;y++ {
    for x := 0;x<=maxX;x++ {
      var p1 Punto
      p1.x=x
      p1.y=y
      if _,ok:=punti[p1]; ok{
        fmt.Print("*")
      }else{
        fmt.Print(" ")
      }
    }
    fmt.Println()
  }

}
