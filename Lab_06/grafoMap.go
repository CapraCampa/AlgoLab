package main
import (
	"fmt"
	"strings"
	"bufio"
	"os"
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
	vertici map[string]* vertice
	adiacenti map[ vertice ][]* vertice
}

func graphNew(n int) *grafo{
	vertici:=make(map[string]*vertice)
	adiacenti:=make(map[vertice][]*vertice)
	for i := 0; i < n; i++ {
		map[{{0,&make([]string)}, ""}]=make([]*vertice,n)
	}
	return &grafo{vertici,adiacenti}
}

func main(){
	scanner:=bufio.NewScanner(os.Stdin)
	var n int
	fmt.Printf("Inserire utenti nel formato 'Nome utente che segue' 'Nome utente seguito' :\n")
	grafo:=graphNew(n)
	for scanner.Scan(){
		parole:=strings.Fields(scanner.Text())
		v1,_:=strconv.Atoi(parole[0])
		v2,_:=strconv.Atoi(parole[1])
		addNode(v2,grafo.adiacenti[v1])
	}
}