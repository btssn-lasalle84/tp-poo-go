package main

import "testing"

func TestGetSoldeCompte(t *testing.T) {
    compte1 := CreerCompte("Bill Gates", 105300000000, "USD");
    var attendu float64 = 105300000000
    obtenu := compte1.GetSolde()
    if obtenu != attendu {
        t.Errorf("obtenu : '%f', attendu : '%f'", obtenu, attendu)
    }
    compte2 := CreerCompte("Moi", 42, "EUR");
    attendu = 42
    obtenu = compte2.GetSolde()
    if obtenu != attendu {
        t.Errorf("obtenu : '%f', attendu : '%f'", obtenu, attendu)
    }
    compte3 := CreerCompte("Maria", 1000, "YEN");
    attendu = 1000
    obtenu = compte3.GetSolde()
    if obtenu != attendu {
        t.Errorf("obtenu : '%f', attendu : '%f'", obtenu, attendu)
    }
}

func TestGetTitulaireCompte(t *testing.T) {
    compte1 := CreerCompte("Bill Gates", 105300000000, "USD");
    attendu := "Bill Gates"
    obtenu := compte1.GetTitulaire()
    if obtenu != attendu {
        t.Errorf("obtenu : '%s', attendu : '%s'", obtenu, attendu)
    }
    compte2 := CreerCompte("Moi", 42, "EUR");
    attendu = "Moi"
    obtenu = compte2.GetTitulaire()
    if obtenu != attendu {
        t.Errorf("obtenu : '%s', attendu : '%s'", obtenu, attendu)
    }
    compte3 := CreerCompte("Maria", 1000, "YEN");
    attendu = "Maria"
    obtenu = compte3.GetTitulaire()
    if obtenu != attendu {
        t.Errorf("obtenu : '%s', attendu : '%s'", obtenu, attendu)
    }
}

func TestGetDeviseCompte(t *testing.T) {
    compte1 := CreerCompte("Bill Gates", 105300000000, "USD");
    attendu := "USD"
    obtenu := compte1.GetDevise()
    if obtenu != attendu {
        t.Errorf("obtenu : '%s', attendu : '%s'", obtenu, attendu)
    }
    compte2 := CreerCompte("Moi", 42, "EUR");
    attendu = "EUR"
    obtenu = compte2.GetDevise()
    if obtenu != attendu {
        t.Errorf("obtenu : '%s', attendu : '%s'", obtenu, attendu)
    }
    compte3 := CreerCompte("Maria", 1000, "YEN");
    attendu = "YEN"
    obtenu = compte3.GetDevise()
    if obtenu != attendu {
        t.Errorf("obtenu : '%s', attendu : '%s'", obtenu, attendu)
    }
}
