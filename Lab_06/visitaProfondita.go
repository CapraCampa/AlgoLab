package main
import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"strconv"
)

type vertice struct {
	valore item
	chiave string
}

type item struct {
	età int
	hobbies *[]string
}

type grafo struct {
	vertici map[string]*vertice
	adiacenti map[vertice][]*vertice
}

func graphNew(n int) *grafo{
	vertici:=make(map[string]*vertice)
	adiacenti:=make(map[vertice][]*vertice)
	return &grafo{vertici,adiacenti}
}

func addNode(riga []string,  g *grafo){
	seguito:=riga[0]
	riga[1]=strings.TrimPrefix(riga[1]," ")
	info:=strings.Split(riga[1], " ")
	var per item
	n,_:=strconv.Atoi(info[0])
	per.età=n
	hobbies:=make([]string,0)
	for i:=1; i<len(info); i++{
		hobbies=append(hobbies,info[i])
	}
	per.hobbies=&hobbies
	(*g).vertici[seguito]=&vertice{per,seguito}
}

func addArc(segue string, seguito string, g *grafo){
	follower:=*(*g).vertici[segue]
	(*g).adiacenti[follower]=append((*g).adiacenti[follower], (*g).vertici[seguito])
}

func leggiGrafo(f *os.File)*grafo{
	scanner:=bufio.NewScanner(f)
	var n int
	g:=graphNew(n)
	for scanner.Scan(){
		riga:=scanner.Text()
		if len(strings.Fields(riga))>2{
			addNode(strings.Split(riga,":"), g)
		}else{
			parole:=strings.Fields(riga)
			addArc(parole[0],parole[1],g)
		}

	}
	return g
}

func stampa(g *grafo){
	fmt.Println("Persone iscritte al social: ")
		for _,slice:=range (*g).vertici{
			fmt.Print((*slice).chiave)
			fmt.Println(" ",(*slice).valore.età, " ")
			hobby(*slice, g)
		}
		fmt.Println()
		for key,slice:=range (*g).adiacenti{
			fmt.Print(key.chiave, " segue:")
			for _,n:=range slice{
				fmt.Print(" ",(*n).chiave)
			}
			fmt.Println()
		}
		fmt.Println()	
}

//Scrivete una funzione che data una stringa A stampi gli hobby dell’utente di nome A e
//l’elenco di tutti gli hobby delle persone che seguono A.
func hobby(verticeA vertice, g *grafo){
	fmt.Print("Hobbbies di ",verticeA.chiave,":" )
	ho:=*verticeA.valore.hobbies
	for _,h:=range ho{
		fmt.Print(" ",h)
	}
	fmt.Println()
}

func hobbySeguentiA(a string, g *grafo){
	verticeA:=*(*g).vertici[a]
	hobby(verticeA,g)
	fmt.Print("Hobbies delle persone che seguono ",a, ":\n")
	for verticeB,slice:=range (*g).adiacenti{
		for _,x:=range slice{
			if *x==verticeA{
				hobby(verticeB,g)
				break
			}		
		}
	}
	fmt.Println()
}

func hobbySeguitiDaA(a string, g *grafo){
	verticeA:=*(*g).vertici[a]
	hobby(verticeA,g)
	fmt.Println("Hobbies delle persone che ",a, " segue:")
	for _,verticeB:=range (*g).adiacenti[verticeA]{
		hobby(*verticeB,g)
	}
	fmt.Println()
}

func dfs1 (g grafo , v string, aux map[string]bool) {
	fmt.Println(v)
	aux[v] = true
	ver:=g.vertici[v]
	for _ , ver2 := range g.adiacenti[*ver] {
		v2:=(*ver2).chiave
		if aux[v2] != true {
			dfs1(g , v2 , aux )
		}
	}
}


func dfs1Pila(g grafo , v string, aux map[string]bool) {
	coda := []string{v}
	aux [v] = true
	for len( coda ) > 0 {
		v := coda [len(coda)-1]
		coda = coda [:len(coda)-1]
		fmt .Println("\t", v)
		ver:=g.vertici[v]
		for _, ver2 :=range g.adiacenti[*ver]{
			v2:=(*ver2).chiave
			if ! aux [ v2 ] {
				coda = append( coda , v2 )
				aux [ v2 ] = true
			}
		}
	}
}


func bfs1 (g grafo , v string, aux map[string]bool) {
	coda := []string{v}
	aux [v] = true
	for len( coda ) > 0 {
		v := coda [0]
		coda = coda [1:]
		fmt .Println("\t", v)
		ver:=g.vertici[v]
		for _, ver2 :=range g.adiacenti[*ver]{
			v2:=(*ver2).chiave
			if ! aux [ v2 ] {
				coda = append( coda , v2 )
				aux [ v2 ] = true
			}
		}
	}
}


func main(){
	f,err:=os.Open("twitter.txt")
	if err!=nil{
		return
	}
	g:=leggiGrafo(f)
	stampa(g)
	hobbySeguitiDaA("valeria",g)
	hobbySeguentiA("riccardo",g)
	aux:=make(map[string]bool)
	fmt.Println("Visita in profondità:")
	dfs1(*g,"valeria",aux)
	aux=make(map[string]bool)
	fmt.Println("Visita in profondità con altro criterio:")
	dfs1Pila(*g,"valeria",aux)
	aux=make(map[string]bool)
	fmt.Println("Visita in ampiezza:")
	bfs1(*g,"valeria",aux)
}