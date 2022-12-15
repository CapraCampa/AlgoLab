package main
import(
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

//funzione di transizione che mi dice data una posizione la posizione adiacente migliore
func mossa(n int, caselle *[]int,passi *[]int)int{
	var max,dado int
	for i:=1; i<=6; i++{
		if n+i>=len(*caselle)-1{
			*passi=append(*passi,i)
			return n+i
		}
		if (*caselle)[n+i]==0{
			(*caselle)[n+i]=n+i+1
		}
		if (*caselle)[n+i]-1>=max{
			dado=i
			max=(*caselle)[n+i]-1
		}
	}
	*passi=append(*passi,dado)
	return max
}

func main(){
	f,err:=os.Open("InputS1.txt")
	if err!=nil{
		return
	}
	scanner:=bufio.NewScanner(f)
	scanner.Scan()
	riga:=strings.Fields(scanner.Text())
	r,_:=strconv.Atoi(riga[0])
	c,_:=strconv.Atoi(riga[1])
	n:=r*c
	caselle:=make([]int,n)
	for scanner.Scan(){
		salto:=strings.Fields(scanner.Text())
		from,_:=strconv.Atoi(salto[0])
		to,_:=strconv.Atoi(salto[1])
		caselle[from-1]=to
	}
	passi:=make([]int,0)
	ultimoPasso:=0
	for{
		if ultimoPasso>=len(caselle)-1{
			break
		}
		ultimoPasso=mossa(ultimoPasso,&caselle,&passi)
	}
	fmt.Printf("Per vincere la partita servono almeno %d lanci\n",len(passi))
	fmt.Println("Un esempio di sequenza di passi è:",passi)
}