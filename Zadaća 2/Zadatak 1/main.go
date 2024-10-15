
/*1. Napiši strukturu koja predstavlja pravokutnik i sadrži polja za širinu i visinu. Kreiraj metode
 za izračunavanje površine i opsega pravokutnika, metode moraju biti vezane za novo kreiranu strukturu
  Pravokutnika. U main funkciji inicijalizirajte jedan pravokutnik, te pozovite iznad kreirane metode za 
  računanje površine i opsega.
*/
package main

import "fmt"

type Pravokutnik struct {
	Visina int
	Širina int
}

func (p Pravokutnik) Površina() int {
	return p.Visina * p.Širina
}

func (p Pravokutnik) Opseg() int{
	return 2 * (p.Visina + p.Širina)
}
func main(){

	p :=Pravokutnik{
		Visina: 24,
		Širina: 12,
	}

		fmt.Println("Visina pravokutnika: ",p.Visina)
		fmt.Println("Širina pravokutnika: ",p.Širina)
		fmt.Println("Površina pravokutnika: ",p.Površina())
		fmt.Println("Opseg pravokutnika: ", p.Opseg())
	



}