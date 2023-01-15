package main

import (
	"fmt"
)

func stampaSubordinati(sub map[string][]string, d string){
	fmt.Print("I subordinati di ",d," sono: ")
	for _,s:=range sub[d]{
		fmt.Print(s, " ")
	}
	fmt.Println()
}

func supervisore(sup map[string]string, s string)string{
	return sup[s]
}

func quantiSenzaSubordinati(sub map[string][]string) int{
	c:=0
	for s,_:=range sub{
		if len(sub[s])==0{
			c++
		}
	}
	return c
}
func stampaImpiegatiSopra(sup map[string]string, s string){
	fmt.Print("Gli impiegati sopra ",s," sono: ")
	for sup[s]!=""{
		fmt.Print(sup[s], " ")
		s=sup[s]
	}
	fmt.Println()
}

func stampaImpiegatiConSupervisore(sup map[string]string){
	fmt.Print("Impiegati con un superivisore: ")
	for _,s:=range sup{
		if len(sup[s])!=0{
			fmt.Print(s," ")
		}
	}
	fmt.Println()
}

func stampaPerLivello(sub map[string][]string, capi []string){
	coda:=make([]string,0)
	fmt.Println("Tutti i dipendenti:")
	for _,s:=range capi{
		coda=append(coda,s)
	}
	for len(coda)!=0{
		el:=coda[0]
		fmt.Print(el, " ")
		coda=coda[1:]
		if len(sub[el])!=0{
			coda=append(coda, sub[el]...)
		}
	}
	fmt.Println()
}

func aggiungi(sup map[string]string, sub map[string][]string, s string, figli []string, capi *[]string){
	if _,ok:=sup[s]; !ok{
		sup[s]=""
		*capi=append(*capi,s)
	}
	for _,f:=range figli{
		sup[f]=s
		if _,ok:=sub[f]; !ok{
			sub[f]=[]string{}
		}
	}
	sub[s]=figli
}

func main(){
	superiori:=make(map[string]string)
	subordinati:=make(map[string][]string)
	capi:=make([]string,0)
	slice:=[]string{"B","C","D"}
	aggiungi(superiori,subordinati,"A",slice,&capi)
	slice=[]string{"E","F"}
	aggiungi(superiori,subordinati,"B",slice,&capi)
	slice=[]string{"H"}
	aggiungi(superiori,subordinati,"G",slice,&capi)
	slice=[]string{"I"}
	aggiungi(superiori,subordinati,"F",slice,&capi)

	stampaSubordinati(subordinati, "A")
	fmt.Println("Il supervisore di H è",supervisore(superiori,"H"))
	stampaImpiegatiSopra(superiori, "I")
	fmt.Printf("Ci sono %d dipendenti senza subordinati\n",quantiSenzaSubordinati(subordinati))
	stampaImpiegatiConSupervisore(superiori)
	stampaPerLivello(subordinati,capi)
}