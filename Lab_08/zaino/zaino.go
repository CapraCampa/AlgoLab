package main

import(
	"fmt"
	"os"
	"bufio"
)

type Ogg struct{
	p int
	v int
}

func bubbleSort(s []Ogg){
	for i:=0; i<(len(s)-i); i++{
		for j:=i+1; j<len(s)-i; j++{
			if s[i].p>s[j].p{
				s[i],s[j]=s[j],s[i]
			}
		}
	}
}



func scheduling(o []Ogg, cap int, n int)int{
	var val [][]int
	for i:=0; i<=n ; i++{
		riga:=make([]int, cap+1)
		for j:=0; j<=cap; j++{
			riga[j]=0
		}
		val=append(val,riga)
	}
	for i:=1; i<=n; i++{
		for j:=0; j<=cap; j++{
			val[i][j]=val[i-1][j]

			if ((j>=o[i-1].p) && (val[i][j] < val[i-1][j-o[i-1].p] + o[i-1].v)){
				val[i][j]=val[i-1][j-o[i-1].p] + o[i-1].v
			}
		}
	}

	num:=n
	peso:=cap
	/*for n!=0{
		if val[n][cap] != val[n-1][cap]{
			//n= oggetto che ho preso
			cap=cap-o[n-1].p
		}
		n--
	}*/
	return val[num][peso]
}


func main(){
	t,err:=os.Open("input.txt")
	if err!=nil{
		return
	}
	oggetti,cap:=leggiInput(t)
	bubbleSort(oggetti)
	max:=scheduling(oggetti,cap, len(oggetti))
	fmt.Println(max)
}

func leggiInput(t *os.File)([]Ogg,int){
	var n, cap, p, v,i int
	scanner:=bufio.NewScanner(t)
	scanner.Scan()
	fmt.Sscanf(scanner.Text(),"%d %d",&n,&cap)
	oggetti:=make([]Ogg,n)
	for scanner.Scan(){
		fmt.Sscanf(scanner.Text(),"%d %d",&p,&v)
		oggetti[i]=Ogg{p,v}
		i++
	}
	return oggetti,cap
}