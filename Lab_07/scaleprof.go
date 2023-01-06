package main

import "fmt"

type board struct {
	n     int
	jumps map[int]int // jumps[i] è la casella di destinazione se ci si trova in i
}

func setBoard() board {
	var r, c int
	fmt.Scan(&r, &c)
	n := r * c

	jumps := make(map[int]int)
	for {
		var start, end int
		_, err := fmt.Scan(&start, &end)
		if err != nil {
			break
		}
		jumps[start] = end
	}
	fmt.Println(jumps)
	return board{n, jumps}

}

func move(b board, start int, dice int) (end int) {
	if el, ok := b.jumps[start+dice]; ok {
		return el
	} else {
		return start + dice
	}
}

// calcola il minimo numero di lanci necessari per arrivare alla casella di valore v partendo dalla casella di valore v0
// modello l'evoluzione del gioco con un grafo (i nodi rappresentano le configurazioni, gli archi le mosse)
// uso una visita in ampiezza del grafo (serve una coda)
// Restituisce 0 se v e v0 coincidono; restituisce -1 se partendo da v0 non si può raggiungere v
// se seq e' pari a 1, stampa anche una sequenza di lanci di dadi di lunghezza minima (se c'è)

func bfs(b board, v0, v int, seq bool) {
	if v0 == v {
		fmt.Println(0)
		return
	}
	prec := make(map[int]int) // per ogni casella, indica la casella precedente
	dice := make(map[int]int) // per ogni casella, indica con che dado ci si è arrivati
	aux := make(map[int]int)  // caselle già visitate, con distanza da v0

	fmt.Println("\naux", aux)

	coda := []int{v0} // all'inizio la coda contiene solo la casella di partenza
	curr := v0
	aux[v0] = 0
	dice[v0] = -1
	fmt.Println("coda", coda, "\naux", aux)

	for len(coda) > 0 {
		curr = coda[0]
		coda = coda[1:]
		fmt.Println("\tposizione", curr)
		//	fmt.Println("coda", coda, "\naux", aux)

		// metti nella coda tutti i vicini non ancora visitati:
		for i := 1; i <= 6; i++ {
			end := move(b, curr, i)
			if aux[end] >= 1 {
				continue
			}

			if _, ok := aux[end]; !ok {
				coda = append(coda, end)
			}

			fmt.Println("end", end, "curr", curr, "i", i)
			prec[end] = curr
			dice[end] = i
			aux[end] = aux[curr] + 1

			if end == b.n {
				fmt.Println("bastano", aux[curr]+1, "mosse")
				if seq {

					printSeq(end, v0, prec, dice)
				}
				return
			}
			//fmt.Println("pos:", curr, "dado", i, "end:", end)

		}
		fmt.Println("coda", coda, "\naux", aux)
		/*var line string
		fmt.Scan(&line)*/

	}

}

func printSeq(v, v0 int, prec map[int]int, dice map[int]int) {
	//printCells(prec)
	for v != v0 {
		fmt.Println(v, "venendo da", prec[v], "con tiro", dice[v])
		v = prec[v]
	}

}

func main() {
	b := setBoard()
	fmt.Println(b.n, b.jumps)

	fmt.Println(move(b, 3, 3))
	bfs(b, 1, 30, true)

}
