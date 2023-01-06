package main
import(
	"fmt"
	"strings"
	"bufio"
	"os"
	"strconv"
)

type info struct{
	prec int //numero di casella precedente
	dado int	//dado con cui è arrivato lì
	lanci int //numero di lanci per arrivarci
}

func move( curr int,  i int,  salti map[int]int)int{
	if arrivo,ok:=salti[curr+i]; ok{
		return arrivo
	}
	return curr+i
}

func stampa( curr int,  visitati map[int]info){
	fmt.Printf("Si può raggiungere %d tramite %d lanci:\n",curr,visitati[curr].lanci )
	var s string = strconv.Itoa(visitati[curr].dado)
	for{
		curr=visitati[curr].prec
		if curr==1{
			break
		}
		s=strconv.Itoa(visitati[curr].dado)+","+s
	}
	fmt.Println(s)
}

func main(){
	f,err:=os.Open("InputS3.txt")
	if err!=nil{
		return
	}
	scanner:=bufio.NewScanner(f)
	scanner.Scan()
	riga:=strings.Fields(scanner.Text())
	r,_:=strconv.Atoi(riga[0])
	c,_:=strconv.Atoi(riga[1])
	ultimaCas:=r*c
	salti:=make(map[int]int)
	for scanner.Scan(){
		salto:=strings.Fields(scanner.Text())
		from,_:=strconv.Atoi(salto[0])
		to,_:=strconv.Atoi(salto[1])
		salti[from]=to
	}
	//fmt.Println(ultimaCas,salti)
	coda:=make([]int,0)
	coda=append(coda,1)
	visitati:=make(map[int]info)
	var curr,next int
	visitati[1]=info{0,0,0}
	loop:
	for len(coda)!=0{
		//fmt.Println("coda:",coda)
		//fmt.Println("visitati:",visitati)
		curr=coda[0]
		coda=coda[1:]
		for i:=1; i<=6; i++{
			next=move(curr,i,salti)
			//fmt.Println("Prossima casella:", next)
			if _,ok:=visitati[next]; !ok{
				coda=append(coda,next)
				visitati[next]=info{curr,i,(visitati[curr].lanci)+1}
			}
			//fmt.Println(next,"e", visitati[next])
			if next>=ultimaCas{
				break loop
			}
		}
	}
	if _,ok:=visitati[ultimaCas]; ok{
				fmt.Println("Sono arrivato alla fine")
				stampa(ultimaCas,visitati)	
	}else{
		fmt.Println("Non posso arrivare alla fine")
	}
}