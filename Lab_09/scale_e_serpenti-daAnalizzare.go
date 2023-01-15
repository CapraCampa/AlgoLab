package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
- E' possibile utilizzare i grafi per modellare il gioco "scale e serpenti"
- I nodi possono rappresentare la posizione della pedina, mentre un arco che
collega due vertici v1 e v2, se è possibile raggiungere v2 partendo da v1.
- Il grafo sarà orientato, dal momento che non è possibile tornare alla casella
precedente una volta fatta una mossa, a meno che non faccio una mossa e
finisco in una casella con un serpente, che mi riporta alla casella precedente
(ma questo è un caso raro).
- Il grafo generalmente non è connesso, ad eccezione di casi eccezionali. Ad
esempio se in ogni casella c'è un serpente che riporta alla casella precedente,
il grafo risultante sarà connesso.
- Il grafo non è pesato, in quanto ci basta sapere da una casella (nodo), quale
altra casella è possibile raggiungere

Approccio al problema:
- Per memorizzare i dati scelgo di utilizzare una mappa. La preferisco rispetto
ad un array in quanto ci evita di utilizzare più memoria (non dobbiamo salvare
tutti gli 0). Se una cella non ha nè basi di scale, nè bocche di serpente, il suo
numero non sarà presente come chiave nella mappa, altrimenti la mappa avrà una
chiave pari al numero della casella e il suo valore corrisponde al numero della
casella al quale si deve andare se si finisce sopra.
- Decido di sfruttare il concetto di pila. In un ciclo continuo ad iterare
finchè la pila non è vuota. Ad ogni iterazione aggiungo alla pila tutte le
possibili mosse che è possibile fare da una certa casella. Alla prossima
iterazione preleverò l'ultima mossa salvata, e la eseguo.
- Per evitare di andare avanti all'infinito, tengo traccia su quali caselle sono
già stato. In quanto stiamo cercando il percorso più breve, è inutile passare
per una casella già visitata in precedenza. In questo caso non aggiungo nessuna
mossa alla pila.
- Nella pila salvo una struct di tipo T che contiene una slice contenente le mosse
dei dadi effettuate, e un intero che indica a che casella siamo arrivati. Il valore
size(slice_mosse) indica il numero di mosse compiute fino a quel momento.
- Quando l'intero all'interno della struct è pari ad n (con n pari al numero di
caselle) salvo l'intera struct in una slice di struct T. Se la slice (all'interno
della struct) che sto per inserire ha size minore di tutte quelle già salvate,
cancello tutte quelle presenti e salvo quella nuova. Se invece il size della slice
che sto per aggiungere è maggiore del size di quelle presenti non la salvo.
- Quando la pila è vuota mi sono rimasti nell'array tutte le posibili combinazioni
di lanci più brevi per quella tabella di gioco.
*/

type Move struct {
	moves []int
	cell  int
}

func main() {
	check_for_stairs_and_snake := true
	r, c, cells := read_input()
	n := r * c

	stack := make([]Move, 0)
	already_visited := make([]bool, n)
	var finish_moves []int

	for i := 0; i < n; i++ {
		already_visited[i] = false
	}

	stack = append(stack, Move{moves: make([]int, 0), cell: 1})

	for len(stack) != 0 {
		fmt.Println(stack)
		move := stack[0]
		stack = stack[1:]

		if move.cell == n && (/*len(move.moves) <= len(finish_moves) ||*/ len(finish_moves) == 0) {
			finish_moves = move.moves
		}

		value, exists := cells[move.cell]
		if exists && check_for_stairs_and_snake {
			if (!already_visited[value-1]) || (value == n) {
				stack = append(stack, Move{moves: move.moves, cell: value})
				already_visited[value-1] = true
			}
		} else {
			for i := move.cell + 1; (i < move.cell+7) && (i <= n); i++ {
				if (!already_visited[i-1]) || (i == n) {
					stack = append(stack, Move{moves: append(move.moves, i-move.cell), cell: i})
					already_visited[i-1] = true
				}
			}
		}
	}

	fmt.Print("Puoi arrivare alla casella finale con ")
	fmt.Print(len(finish_moves))
	fmt.Println(" mosse e con i seguenti lanci:")

	for i := 0; i < len(finish_moves); i++ {
		fmt.Print(finish_moves[i])
		fmt.Print(" ")
	}

	fmt.Println()
}

func read_input() (r int, c int, cells map[int]int) {
	var c1, c2 int

	file, err := os.Open("input.txt")

	if err == nil {
		scanner := bufio.NewScanner(file)

		if scanner.Scan() {
			line := scanner.Text()

			fmt.Sscanf(line, "%d %d", &r, &c)
			cells = make(map[int]int)
		}

		for scanner.Scan() {
			fmt.Sscanf(scanner.Text(), "%d %d", &c1, &c2)
			cells[c1] = c2
		}
	}

	return
}

/*
Traccia per l'analisi:

Considerate l’esercizio “scale e serpenti” della scheda su “modellazione con grafi” e lo svolgimento proposto in questo programma.

Leggete innanzitutto il commento e valutate quanto segue:
1. La descrizione del grafo è chiara e corretta? In caso contrario, come la correggereste/chiarireste?
La descrizione del grafo è completa, è indicato correttamente cosa archi e vertici rappresentano, sono indicate le varie caratteristiche del grafo e sono argomentate, anche se alcune argomentazioni non risultano del tutto corrette

2. Il grafo descritto modella appropriatamente la situazione descritta nella traccia dell’esercizio?
Il grafo scelto è abbastanza corretto, la scelta dei vertici e archi è appropriata così come l'idea di un grafo orientato e non pesato per poter risolvere il problema posto nel testo.
Tuttavia in base alla scelta dei vertici ed archi e in base alla descrizione dell'esercizio il grafo dovrebbe risultarre connesso, infatti ogni vertice è collegato a un altro poiché ogni casella è collegata a un'altra tramite un lancio di dado o una scala/serpente.

3. Il commento contiene una formulazione del problema da risolvere il termini di grafi? Se sì, quale è questa formulazione? Se no, come lo formulereste?
No, il commento non contiene una formulazione del problema in termini di grafi.
Dato il grafo precedentemente descritto per soddisfare la richiesta bisogna risolvere un problema di ricerca del cammino minimo tra due vertici (il vertice che rappresenta la configurazione in cui la pedina si trova sulla casella iniziale 1 e il vertice in cui la pedina si trova sulla casella finale n)

4. Il commento contiene una descrizione dell’algoritmo per risolvere il problema (poi implementato nel programma)? Se sì, dove si trova questa descrizione?
Nel commento viene descritto in parte l'algoritmo nelle righe in cui si parla della pila e delle iterazioni successive per trovare il minimo numero di lanci di dadi (da riga 30 a riga 47). Tuttavia vi sono alcuni errori nella descrizione del procedimento e delle strutture utilizzate.

Considerate ora il codice:
1. Come sono rappresentati nel programma i vertici e gli archi del grafo? Sono rappresentazioni implicite o esplicite?
I vertici sono rappresentati con numeri da 1 a n (con n uguale al numero di caselle), non contengono altri valori e per questo non è necessario creare struct o altre strutture particolari. Alcuni archi sono indicati tramite una mappa che agisce come una sorta di lista di adiacenza, se in una casella inizia una scala o un serpente usa il numero di quella casella come chiave e il numero della casella d'arrivo come valore associato.
Non salva nella mappa tutti gli archi tra i vertici perché sfrutta una proprietà dell'esercizio per cui osservando un vertice si può già sapere a quali vertici è collegato in uscita (infatti da ogni casella posso sempre calcolare le caselle collegate poiché so dal testo che da ogni casella posso spostarmi di un numero variabile di posizioni, che vanno da 1 a 7, fino a un'altra casella, con le eccezioni delle scale e dei serpenti individuate tramite la mappa)

2. Cosa fa il ciclo for di riga 85-89? Provate a descriverlo facendo riferimento al grafo usato per modellare l’evoluzione del gioco.
Dato il vertice in cui mi trovo, sapendo che non è collegato tramite scale/serpenti a un altro vertice, controllo tutti i vertici a cui è collegato in uscita e se trovo un vertice che non ho ancora visitato o se si tratta del vertice d'arrivo lo aggiungo all'elenco di vertici del percorso attuale e lo aggiungo allo "stack" 

3. Considerate la variabile “stack”. Che tipo di struttura dati implementa?
Implementa una coda in cui il primo elemento prelevato è il primo inserito (struttura Last In First Out).
Pertanto il nome "stack" risulta confusionario, visto che lo stack è una struttura di tipo First In First Out.

4. A cosa serve il controllo di riga 74? Vi pare necessario? Giustificate la risposta
Serve per salvare la sequenza migliore di mosse una volta arrivato all'ultima casella, la salvo se non ne ho ancora salvata nessuna o se è di lunghezza minore o uguale di quella già presente. Per come è implementato l'esercizio non mi sembra necessario visto che appena trovo un percorso per arrivare all'ultima casella so già che tutti i percorsi che troverò in seguito saranno di lunghezza uguale o maggiore e il problema mi chiede uno qualsiasi dei possibili percorsi minimi

5. Se avevate risposto no all’ultima domanda a proposito del commento, date voi una descrizione dell’algoritmo implementato per risolvere il problema (è utile descriverlo facendo riferimento al grafo).
Per risolvere questo problema viene inizializzata una coda (lo "stack") nella quale inserirò delle strutture composte da un numero che indica il vertice in cui si è arrivati e dall'elenco di lanci di dadi utilizzati per arrivare al vertice corrente partendo dal vertice iniziale.
All'inizio la coda contiene solo una struttura composta dal vertice iniziale e un elenco vuoto di lanci per arrivarci.
Ad ogni iterazione prelevo la prima struttura della coda e visito i vertici raggiungibili dal vertice V della struttura che ho prelevato, inserendo nella coda tutte le strutture composte da un vertice che non ho ancora visitato V1 e l'elenco di mosse per arrivare al vertice V con in aggiunta l'ultima mossa per passare dal vertice V al vertice V1.
Ogni volta che arrivo al vertice finale verifico se ci sono arrivato per la prima volta o se ho utilizzato un numero di mosse inferiore o uguale al numero minimo di mosse già individuato, se mi trovo in uno di questi due casi salvo l'elenco di mosse utlizzate in un'altra variabile.  
Proseguo queste iterazioni finché la coda non è vuota, ottenendo nella variabile "finish_moves" l'ultima sequenza di mosse individuata di lunghezza minima.

6. Vi sembra che questo algoritmo risolva correttamente il problema? Giustificate la risposta (di nuovo, è utile fare riferimento al grafo).
Questo algoritmo risolve correttamente il problema, la richiesta di trovare uno dei percorsi minimi per andare dalla casella iniziale alla casella finale è infatti soddisfatta, così come alcune richieste secondarie del testo del problema (come la modifica del codice affinché sia in grado di risolvere lo stesso problema ma senza utilizzare le scale o i serpenti).


Avete altri commenti o osservazioni relativi allo svolgimento di questo esercizio?
L'algoritmo utilizzato risulta corretto ma non molto efficiente, infatti come si nota anche dal commento precedente al codice l'algoritmo era stato inizialmente pensato per restituire tutte le sequenze di mosse di lunghezza minima per arrivare dal vertice iniziale al vertice finale. Per questo motivo l'algoritmo non si ferma al primo percorso trovato ma continua le iterazioni, effettuando quindi molto più lavoro di quanto sia necessario per risolvere il problema del testo.
Inoltre la descrizione dell'algoritmo nel commento differisce in alcuni aspetti dall'implementazione nel codice. Non solo, per l'appunto, non vengono restituite tutte le sequenze minime ma solo l'ultima (in linea con la richiesta del problema ma non col la descrizione nel commento), ma viene anche descritto l'utilizzo e il funzionamento di una pila quando nel codice viene invece utilizzata una coda. 

*/
