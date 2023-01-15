package main

import (
		"fmt"
)

func quickSort(array []int, i int, f int){
	if f-i<=1{
		return
	}
	m:=partiziona(array,i,f)
	quickSort(array,i,m)
	quickSort(array,m+1,f)
}

func partiziona(array []int, i int, f int) int{
	var sx, dx int
	sx=i
	dx=f-1
	for ; ;{
		for ; sx<dx;{
			if array[dx]<=array[i]{
				break
			}
			dx--
		}
		for ;sx<dx; {
			if array[sx]>array[i]{
				array[sx],array[dx]=array[dx],array[sx]
				break
			}
			sx++
		}
		if sx==dx{
			break
		}
	}
	array[sx],array[i]=array[i],array[sx]
	return sx
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
	fmt.Println(arrayB)
	quickSort(arrayB, 0, m)
	fmt.Println(arrayB)
	
}