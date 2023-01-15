package main

import(
	"fmt"
	"bufio"
	"os"
)

type Evento struct{
	i int
	f int
	v int
}


//Uso selection sort
func ordina(eventi []Evento){
	var min, iMin int
	for j:=0; j<len(eventi); j++{
		min=(eventi[j]).f
		iMin=j
		for k:=j+1; k<len(eventi);k++{
			if min>(eventi[k].f){
				min=eventi[k].f
				iMin=k
			}
		}
		eventi[j], eventi[iMin]=eventi[iMin], eventi[j]
	}
}

func ultimoPrima(e []Evento, in int)int{
	for i:=in-1;i>=0; i--{
		if e[i].f<e[in].i{
			return i
		}
	}
	return -1
}

func opt(e []Evento, in int)int{
	if in<0{
		return 0
	}
	if (opt(e, ultimoPrima(e, in))+e[in].v) < opt(e,in-1){
		return opt(e,in-1)
	}else{
		return (opt(e, ultimoPrima(e, in))+e[in].v)
	}

}

func scheduling(e []Evento, n int)(int, []Evento){
	var max,maxIn int
	valori:=make([]int,n)
	for i:=0; i<len(e); i++{
		valori[i]=opt(e,i)
		//fmt.Println(e[i],valori[i])
		if max<opt(e,i){
			max=opt(e,i)
			maxIn=i
		}
	}
	seq:=[]Evento{e[maxIn]}
	for i:=maxIn; i>0; i--{
		if e[maxIn].v+valori[i-1]==valori[maxIn]{
			seq=append([]Evento{e[i-1]},seq...)
			maxIn=i-1
		}
	}
	return max,seq
}

func main(){
	t,err:=os.Open("input.txt")
	if err!=nil{
		return
	}
	eventi,n:=leggoInput(t)
	ordina(eventi)
	vMax,seq:=scheduling(eventi,n)
	fmt.Println("Massimo valore possibile:",vMax)
	fmt.Println("Possibile sequenza di programmi:",seq)
}

func leggoInput(t *os.File) ([]Evento, int){
	
	scanner:=bufio.NewScanner(t)
	var n,i,f,v,j int
	scanner.Scan()
	fmt.Sscanf(scanner.Text(),"%d", &n)
	eventi:=make([]Evento,n)
	for scanner.Scan(){
		fmt.Sscanf(scanner.Text(),"%d-%d %d", &i, &f, &v)
		eventi[j]=Evento{i,f,v}
		j++
	}
	return eventi, n
}