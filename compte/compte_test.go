package main

import "testing"

func TestCrediterCompte(t *testing.T) {
    compte := CreerCompte("Moi", 0, "EUR");
    var attendu float64 = 42
    compte.Crediter(42)
    obtenu := compte.GetSolde()
    if obtenu != attendu {
        t.Errorf("obtenu : '%f', attendu : '%f'", obtenu, attendu)
    }
}

func TestDebiterCompte(t *testing.T) {
    compte := CreerCompte("Moi", 42, "EUR");
    var attendu float64 = 0
    compte.Debiter(42)
    obtenu := compte.GetSolde()
    if obtenu != attendu {
        t.Errorf("obtenu : '%f', attendu : '%f'", obtenu, attendu)
    }
}

func TestCrediterDebiterCompte(t *testing.T) {
    compte := CreerCompte("Moi", 0, "EUR");
    var attendu float64 = 1
    compte.Crediter(1)
    obtenu := compte.GetSolde()
    if obtenu != attendu {
        t.Errorf("obtenu : '%f', attendu : '%f'", obtenu, attendu)
    }
    attendu = 0
    compte.Debiter(1)
    obtenu = compte.GetSolde()
    if obtenu != attendu {
        t.Errorf("obtenu : '%f', attendu : '%f'", obtenu, attendu)
    }
}