/* MODELLAZIONE
NODI: svincoli delle gallerie
ARCHI: gallerie

c'è un arco tra 2 svincoli se esiste una galleria che li colleghi, e il peso su ogni arco indica la luminosità della galleria

il grafo è non orientato, connesso e pesato

si tratta di un problema di esistenza di un cammino, scegliendo sempre e solo l'arco con peso minore

ALGORITMO:  c'è una lista di nodi visitati, un dato che indica il n di archi attraversati, e un dato che indica in che nodo siamo

    se sono in S finito
	se non sono in S: scelgo l'arco con peso minore e mi muovo nel nodo collegato con quell'arco,
	  	si aggiunge 1 alla variabile del numero di archi attraversati
	se il nodo è nella lista dei visitati: non esiste un cammino che rispetti la strategia
	se il nodo non è nella lista si inserisce il nodo nella lista e si ripete l'algoritmo

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type arcoMigliore []int

//pos 0 = vicino migliore
//pos 1 = luminosità

type Grafo = []arcoMigliore //indice = nodo di partenza

func main() {
	var numGallerie int
	var nodiVisitati []int
	var posizione int

	g, h, s := CreaGrafo()

	posizione = h

	for {
		if posizione == s {
			fmt.Println(numGallerie)
			return
		}
		posizione = g[posizione-1][0]
		numGallerie++
		if !(CercaInSlice(nodiVisitati, posizione)) {
			nodiVisitati = append(nodiVisitati, posizione)
		} else {
			fmt.Println(-1)
			return
		}
	}
}

func CreaGrafo() (g Grafo, h int, s int) { // restituisce il grafo con solo le gallerie sceglibili da harmony, il nodo di partenza dell'algoritmo (Harmony) e il nodo d'arrivo(Sarah)
	var num []int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		campi := strings.Split(scanner.Text(), " ")
		for _, cont := range campi {
			n, _ := strconv.Atoi(cont)
			num = append(num, n)
		}
		if len(num) == 4 { //solo prima riga input
			g = make(Grafo, num[0])
			h = num[2]
			s = num[3]
			num = nil
			continue
		}

		if (g[num[0]-1] == nil) || (g[num[0]-1][1] > num[2]) { //se l'arco che stavo scegliendo è più luminoso scelgo quello nuovo
			g[num[0]-1] = arcoMigliore{num[1], num[2]}
		}
		if (g[num[1]-1] == nil) || (g[num[1]-1][1] > num[2]) { //se l'arco che stavo scegliendo è più luminoso scelgo quello nuovo
			g[num[1]-1] = arcoMigliore{num[0], num[2]}
		}

		num = nil
	}
	//fmt.Println(g)
	return g, h, s
}

func CercaInSlice(s []int, n int) bool {
	for _, num := range s {
		if num == n {
			return true
		}
	}
	return false
}

/*
Traccia per l'analisi:

Considerate l’esercizio “Sunnydale” della scheda su “modellazione con grafi” e lo svolgimento proposto in questo programma.
Rileggete le domande nella sezione 1.2 e valutate se il commento fornisce risposte corrette, chiare e complete
alle quattro domande. In caso contrario, come le correggereste/completereste?
Il commento fornisce risposte perlopiù corrette, anche se non si può affermare che il grafo sia connesso visto che in base a come è scritto il testo potrei anche avere due svincoli collegati tra loro ma scollegati dal resto degli svincoli
Il grafo descritto può essere appropriato per risolvere il problema.

Considerate ora il codice:
1. Il grafo che viene manipolato dal programma corrisponde a quanto descritto nel commento?
    1. Se sì, descrivete come è implementato il grafo (è utile fare riferimento ai nomi dei tipi e delle variabili usate, e alle righe di codice rilevanti)
    2. Se no, descrivete quale grafo è effettivamente implementato nel codice descrivete come è implementato (è utile fare riferimento ai nomi dei tipi e delle variabili usate, e alle righe di codice rilevanti)
No, il grafo descritto e il grafo implementato non corrispondono. Il grafo implementato ha sempre per nodi gli svincoli delle gallerie ma come archi solo le gallerie con luminosità minima. Dato uno svincolo e le varie gallerie che si diramano da esso viene infatti considerato come arco orientato nel grafo solo la galleria di luminosità minore. Il grafo risulta quindi orientato e pesato.
Viene implementato con le variabili a riga 31 e 36, ovvero tramite una slice di slice in cui i nodi sono indicati implicitamente come indici della slice. Ad ogni posizione della slice (che corrisponde a un vertice) viene associata una slice contenente il vertice a cui è collegato il vertice individuato dall'indice e il peso dell'arco tramite cui i due vertici sono collegati.

2. Considerate la definizione del tipo arcoMigliore alla riga 31 e verificate come viene usato nelle righe 85-89. Descrivete a parole questo tipo: cosa rappresentano gli elementi della slice? E’ una definizione appropriata? Perché? La modifichereste? Come?
La slice ArcoMigliore è costituita da 2 numeri, il primo indica il vertice di arrivo dell'arco considerato e il secondo considera il peso di questo arco. Considerando che questa slice per nostra definizione è sempre costituita da due soli elementi sarebbe stato sufficiente implementare questo tipo con una struct con due campi e sarebbe stato anche più comprensibile per la lettura del codice.

3. Considerate l’if nelle righe 77-83: è possibile portarlo fuori dal ciclo? Come? E’ opportuno farlo? Perché?
Sì, avrei potuto effettuare la lettura di una riga prima di iniziare ciclo ed eseguire subito le righe 78-81. Questa soluzione è opportuna perché so dal testo che questa sarà la struttura della sola prima riga di input, non conviene eseguire a ogni iterazione successiva alla prima un "if" in cui so già che non entrerò più.

4. Considerate la slice nodiVisitati definita alla riga 40. Cosa rappresenta e come viene usata? E’ una scelta appropriata? In caso negativo, come la modifichereste?
Si tratta di un insieme contenente tutti i nodi che ho già visitato, ogni volta che ne visito uno nuovo lo aggiungo alla fine della slice. Questa struttura è necessaria per il corretto funzionamento dell'algoritmo ma tuttavia poteva essere implementata in modo più efficiente. 
nfatti se utilizziamo una slice contenente il valore dei nodi già visitati, quando dobbiare verificare se un nodo è già presente dovremo impiegare tempo O(n) (viene implementata una ricerca lineare) o se anche implementassimo un algoritmo di ricerca più efficiente impiegheremmo comunque tempo O(nlgn).
Sfrutto il fatto che ciascun nodo sia rappresentabile implicitamente come indice di una slice per creare una slice di lunghezza n contenente valori booleani.
In questo modo accedo al vertice che mi interessa tramite indice (impiegando tempo costante O(1)) e posso subito sapere se è già visitato (trovo valore true) o meno (trovo valore false). 

Avete altri commenti o osservazioni relativi allo svolgimento di questo esercizio?

*/
