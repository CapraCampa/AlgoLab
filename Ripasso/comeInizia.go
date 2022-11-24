package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
  var quanteA,quanteB int
    scanner:=bufio.NewScanner(os.Stdin)
    for i := 0;i < 10;i++ {
      scanner.Scan()
      primaLett:=[]rune(scanner.Text())
      if primaLett[0]=='a'{
        quanteA++
      }else if primaLett[0]=='b'{
        quanteB++
      }
    }
    fmt.Printf("Frasi che iniziano con a: %d \nFrasi che iniziano con b: %d\n", quanteA, quanteB)
}
