package main

import (
		"fmt"
)

func piuPiccolo(arrayA []int, n int, min int) (int) {
	for i:=0; i<n; i++{
		if arrayA[i]>=min{
			return 0
		}
	}
	return 1
}

func trovaMin(array []int,m int) int{
	min:=array[0]
	for i:=1; i<m; i++{
		if min>array[i]{
			min=array[i]
		}
	}
	return min
}

func main(){
	var n, m,num int
	fmt.Scan(&n)
	fmt.Scan(&m)
	arrayA:=make([]int,n)
	arrayB:=make([]int,m)
	for i:=0; i<n; i++{
		fmt.Scan(&num)
		arrayA[i]=num
	}
	for i:=0; i<m; i++{
		fmt.Scan(&num)
		arrayB[i]=num
	}
	min:=trovaMin(arrayB,m)
	if piuPiccolo(arrayA,n,min)==0{
		fmt.Println("Esiste almeno un elemento in B minore o uguale di un elemento di A")
	}else{
		fmt.Println("Ogni elemento di A è strettamente minore di ogni elemento di B")
	}

}