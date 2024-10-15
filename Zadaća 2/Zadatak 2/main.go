/*2. Napiši strukturu koja predstavlja adresu, koja sadrži polja za grad i ulicu.
Zatim napiši strukturu koja predstavlja osobu, koja sadrži ime, godine i adresu.
Kreiraj metodu koja ispisuje puni opis osobe, uključujući njezinu adresu. (Sva polja) 
u formatu: Ime Prezime, 20 godina, živi u Grad, Ulica.
*/

package main

import "fmt"

type Adresa struct{
	Grad string
	Ulica string
}

type Osoba struct{
	Ime string 
	Godina int 
	Adresa Adresa
}

func (o Osoba) Podaci(){
	 fmt.Printf("%s, %d godina, živi u %s, %s",o.Ime,o.Godina,o.Adresa.Grad,o.Adresa.Ulica)
}

func main(){

	o:=Osoba{
		Ime: "John Doe",
		Godina: 20,
		Adresa: Adresa{
			Grad: "Mostar",
			Ulica: "Splitska bb",
		},
		
	}
	o.Podaci()

}