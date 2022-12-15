package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"os"
)

func main() {
	f, err := os.Open("Input1.txt")
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	parametri:=strings.Fields(scanner.Text())
	n,_:=strconv.Atoi(parametri[0])
	h,_:=strconv.Atoi(parametri[2])
	s,_:=strconv.Atoi(parametri[3])
	g := make([]int, n)
	gp := make([]int, n)

	for scanner.Scan(){
		par:=strings.Fields(scanner.Text())
		n1,_:=strconv.Atoi(par[0])
		n2,_:=strconv.Atoi(par[1])
		p,_:=strconv.Atoi(par[2])
		if g[n1-1] == 0 || p < gp[n1-1] {
			gp[n1-1]=p
			g[n1-1] = n2
		}
		if g[n2-1] == 0 || p < gp[n2-1] {
			gp[n2-1]=p
			g[n2-1] = n1
		}
	}
	fmt.Println(g)
	visti:=make([]bool,n)
	c:=0
	for{
		c++
		visti[h-1]=true
		h=g[h-1]
		if h==s{
			break
		}
		if visti[h-1]==true{
			fmt.Println(-1)
			return
		} 
		
	}
	fmt.Printf("Ci vogliono %d passi",c)
}