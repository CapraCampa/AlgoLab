package main

import (
	"fmt"
)

type dipendente struct {
  supervisore *dipendente
	subordinato *dipendente
	collega     *dipendente
	nome        string
}

type capo struct {
	sub *dipendente
}

func creaDipendente(persona *dipendente, nome string) *dipendente {
	return &dipendente{persona,nil, nil, nome}
}

func cercaInAlbero(nome string, radici []capo) *dipendente {
	for _, persona := range radici {
		p := persona.sub
		for p != nil && p.nome != nome {
			p = p.collega
		}
		if p.nome == nome {
			return p
		}
	}
  return nil
}

func stampaSubordinati(dipendente *dipendente) {
	if dipendente != nil {
    fmt.Println(dipendente.nome)
		for dipendente.subordinato != nil || dipendente.collega!=nil{
      stampaSubordinati(dipendente.collega)
			stampaSubordinati(dipendente.subordinato)
		}
	}
  return
}

func main() {
	radici := make([]capo, 0)
	var primo capo
	primo.sub = creaDipendente(nil,"A")
	primo.sub.subordinato = creaDipendente(primo.sub,"B")
	primo.sub.subordinato.collega = creaDipendente(primo.sub.subordinato,"C")
	primo.sub.subordinato.collega.collega = creaDipendente(primo.sub.subordinato.collega,"D")
	primo.sub.subordinato.subordinato = creaDipendente(primo.sub.subordinato,"E")
	primo.sub.subordinato.subordinato.collega = creaDipendente(primo.sub.subordinato.subordinato,"F")
	primo.sub.subordinato.subordinato.collega.subordinato = creaDipendente(primo.sub.subordinato.subordinato.collega,"I")
	radici = append(radici, primo)
	var secondo capo
	secondo.sub = creaDipendente(nil,"G")
	secondo.sub.subordinato = creaDipendente(secondo.sub,"H")
	radici = append(radici, secondo)
  var nome string
  fmt.Scan(&nome)
	dipendente := cercaInAlbero(nome, radici)
	stampaSubordinati(dipendente)
}
