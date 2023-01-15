// DIPENDENTI

/*
Uso la struttura dati albero.
Non sapendo quanti subordinati ha un dipendente, ho usato una mappa per i subordinati di una struttura dipendente (struttura con nome, e puntatore al superiore)
e una lista contenente tutti i dipendenti in elenco.
*/

package main

import "fmt"

type dipendente struct {
	name      string
	val       int
	superiore *dipendente
}

func main() {

	var subordinati map[dipendente][]dipendente = make(map[dipendente][]dipendente)
	//var superiori map[dipendente]dipendente = make(map[dipendente]dipendente)
	var lista []dipendente

	A := &dipendente{"A", 1, nil}
	B := &dipendente{"B", 3, A}
	C := &dipendente{"C", 12, A}
	D := &dipendente{"D", 30, A}
	E := &dipendente{"E", 9, C}
	F := &dipendente{"F", 87, D}
	G := &dipendente{"G", 54, D}

	subordinati[*A] = append(subordinati[*A], *B, *C, *D)
	subordinati[*C] = append(subordinati[*C], *E)
	subordinati[*D] = append(subordinati[*D], *F, *G)

	lista = append(lista, *A, *B, *C, *D, *E, *F, *G)

	stampaSubordinati(B, subordinati)
	fmt.Println()
	fmt.Println(quantiSenzaSubordinati(lista, subordinati))
}

func stampaSubordinati(dipendente *dipendente, subordinati map[dipendente][]dipendente) {
	for _, subordinato := range subordinati[*dipendente] {
		fmt.Print(subordinato.name, " ")
	}
}

func quantiSenzaSubordinati(lista []dipendente, subordinati map[dipendente][]dipendente) int {
	count := 0
	for _, dipendente := range lista {
		_, haSub := subordinati[dipendente]
		if !haSub {
			count++
		}
	}
	return count
}

func supervisore(dipendente dipendente) {
	fmt.Println(dipendente.superiore)
}

/* Traccia per l'analisi:
Considerate l’esercizio “Dipendenti” della scheda di modellazione con alberi
e lo svolgimento proposto in questo programma.

1. Il comento descrive come si può modellare la situazione (domanda 1, parte 1)?
	* Se sì, la descrizione è corretta, chiara, completa?
	* Se no, considerate quanto scritto nel codice e completate il commento rispondendo alla domanda 1, parte 1.
No, non descrive come ogni informazione viene modellata a livello di struttura dati ed è poco precisa.
Modello la situazione come un albero.
Ogni nodo rappresenta un dipendente dell'azienda e i suoi figli rappresentano i suoi subordinati nell'azienda. Ogni nodo rispetto al figlio corrisponde al supervisore del dipendente rappresentato da quel figlio. la radice è rappresentata da un dipendente di livello massimo, ovvero un dipendente privo di supervisore. Ogni nodo può avere un numero arbitrario di figli (poiché ogni dipendente può avere un numero arbitrario di subordinati).

2. Il commento descrive come è stata implementata la struttura dati scelta (domanda 2)?
	1. Se sì, valutate se la descrizione è coerente con quanto scritto nel codice
	ed eventualmente correggete/chiarite tale descrizione.
    2. Se no, considerate quanto scritto nel codice e descrivete come è stata implementata la struttura dati scelta.
	No, il commento descrive molto poco come è stata implementata la struttura dati, non indica in modo chiaro come ogni parte della struttura sia stata implementata.
	Ogni dipendente è implementato come una struttura che contiene un nome, un valore numerico e un puntatore alla struttura del suo superiore.
	Non potendo sapere in anticipo quanti subordinati ha un dipendente, ho implementato questa relazione usando una mappa che associa a ogni struttura "dipendente" una slice con i suoi dipendenti subordinati, inserisco un nuovo dipendente come chiave nella mappa solo se so che ha almeno un subordinato.
	Creo anche una lista per contenere tutti i dipendenti dell'azienda.

3. La struttura dati scelta modella appropriatamente la situazione descritta nella traccia dell’esercizio?
(se vi è incoerenza tra quanto scritto nel commento e quanto implementato nel codice, fate riferimento al codice).
In parte, utilizzare un albero è una scelta adatta per rappresentare tutte le relazioni richieste dall'esercizio se fosse presente un solo dipendente di livello massimo. Tuttavia per essere una struttura corretta avrei dovuto considerare di avere più alberi, infatti il testo mi dice che possono esserci più dipendenti di massimo livello. Se utilizzo una struttura albero non posso rappresentare i diversi dipendenti di livello massimo.

4. Valutate le scelte relative all’implementazione della struttura dati:
le considerate appropriate per potrer svolgere i compiti in maniera efficiente?
Altrimenti, come le correggereste/migliorereste?
L'implementazione di questa struttura fornisce tutti i dati richiesti e i collegamenti necessari per effettuare le operazioni richieste dell'esercizio.
Tuttavia trovo questa implementazione poco efficiente e poco intuitiva per diversi motivi.
I nomi delle variabili sono poco coerenti ("lista" per indicare un insieme implementato da una slice)-
Occupa più spazio di quanto necessario, se avesse inserito nella mappa tutti i dipendenti non avrebbe avuto bisogno di creare una nuova variabile.
Le informazioni sono inoltre frammentate, se un dipendente ha un subordinato vuol dire che analogamente quel subordinato ha un supervisore, quando modifico la mappa devo anche modificare il puntatore interno alla struct.



Avete altri commenti o osservazioni relativi allo svolgimento di questo esercizio?
*/
