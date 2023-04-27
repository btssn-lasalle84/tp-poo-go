package main

import (
	"fmt"
)

// TODO

func main() {
	compte1 := CreerCompte("Bill Gates", 105300000000, "USD")
	compte1.Afficher()
	compte2 := CreerCompte("Moi", 42, "EUR")
	compte2.Afficher()
	compte3 := CreerCompte("Maria", 1000, "YEN")
	compte3.Afficher()

	compte2.Debiter(42)
	fmt.Printf("Nouveau solde : %f %s\n", compte2.GetSolde(), compte2.GetDevise())
	compte2.Crediter(42)
	fmt.Printf("Nouveau solde : %f %s\n", compte2.GetSolde(), compte2.GetDevise())
}
