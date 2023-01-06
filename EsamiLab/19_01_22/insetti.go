package main
import (
		"fmt"
	)

	/*Tempo impiegato in funzione di n ed m:
		ho un for iniziale per riempire il vettore che impiega tempo O(n)
		ho poi un secondo for per l'elaborazione dei valori che impiega tempo O(n)
		ho zero somme
	 */
func main(){
	var n,m,num int
	fmt.Scan(&n)
	fmt.Scan(&m)
	valori:=make([]int,n)
	for i:=0; i<n; i++{
		fmt.Scan(&num)
		valori[i]=num
	}
	var aumento, aumentoM int
	for i:=0; i<len(valori)-1;i++{
		if valori[i]<valori[i+1]{
			aumento++
		}
		if m==1{
			aumentoM=aumento
			continue
		}else if (i+m)<len(valori) && valori[i]<valori[i+m]{
			aumentoM++
		}
	}
	fmt.Printf("%d %d",aumento,aumentoM)
}