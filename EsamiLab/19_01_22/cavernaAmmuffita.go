package main
import (
	"fmt"
	"math"
)

type vertice struct{
	i int
	j int
}

func findmin(visitati map[vertice]bool, distanze [][]int)vertice{
	min:=math.MaxInt
	var v vertice
	for i:=0; i<len(distanze); i++{
		for j:=0; j<len(distanze[0]);j++{
			if _,ok:=visitati[vertice{i,j}]; !ok && min>distanze[i][j]{
				min=distanze[i][j]
				v.i=i
				v.j=j
			}
		}
	}
	return v
}

func aggiorna(p int, caverna [][]int,z vertice,distanze [][]int){
	x:=z.i
	y:=z.j
	if x+1<len(caverna) && distanze[x+1][y]>distanze[x][y]+caverna[x+1][y]{
		distanze[x+1][y]=distanze[x][y]+caverna[x+1][y]
	}
	if y+1<len(caverna[0]) && distanze[x][y+1]>distanze[x][y]+caverna[x][y+1]{
		distanze[x][y+1]=distanze[x][y]+caverna[x][y+1]
	}
	if p==1{
		if x-1>=0 && distanze[x-1][y]>distanze[x][y]+caverna[x-1][y]{
			distanze[x-1][y]=distanze[x][y]+caverna[x-1][y]
		}
		if y-1>=0 && distanze[x][y-1]>distanze[x][y]+caverna[x][y-1]{
			distanze[x][y-1]=distanze[x][y]+caverna[x][y-1]
		}	
	}
}


func dijkstra(p int, n int, caverna [][]int) int{
	distanze:=make([][]int,n)
	for i:=0; i<n; i++{
		rigaD:=make([]int,n)
		for j:=0; j<n; j++{
			rigaD[j]=math.MaxInt
		}
		distanze[i]=rigaD
	}
	count:=0
	visitati:=make(map[vertice]bool)
	x:=vertice{0,0}
	distanze[0][0]=0
	for count<(n*n){
		x=findmin(visitati,distanze)
		count++
		visitati[x]=true
		aggiorna(p,caverna,x,distanze)
	}
	return distanze[n-1][n-1]
}

func main(){
	var n,num int
	fmt.Scan(&n)
	caverna:=make([][]int,n)
	for i:=0; i<n; i++{
		riga:=make([]int,n)
		for j:=0; j<n; j++{
			fmt.Scan(&num)
			riga[j]=num
		}
		caverna[i]=riga
	}
	fmt.Printf("r1: %d r2: %d\n",dijkstra(1,n,caverna),dijkstra(2,n,caverna))
}