package main

import (
	"testing"
    "fyne.io/fyne/v2/test"
	"os"
)

func TestUICompte(t *testing.T) {
    CreerFenetreCompte()

	if txtTitre.Text != "Compte" {
		t.Errorf("Titre incorrect : '%s', attendu : 'Compte'", txtTitre.Text)
	}

    if txtSolde.Text != FloatToString(monCompte.GetSolde()) {
		t.Errorf("Solde incorrect : '%s', attendu : '%s'", txtSolde.Text, FloatToString(monCompte.GetSolde()))
	}

    if txtMontant.Text != "Montant" {
		t.Errorf("Label incorrect : '%s', attendu : 'Montant'", txtMontant.Text)
	}

    test.Type(edtMontant, "100")
    if edtMontant.Text != "100" {
		t.Errorf("Éditeur incorrect : '%s', attendu : '100'", edtMontant.Text)
	}
}

func TestCrediterCompte(t *testing.T) {
    CreerFenetreCompte()

    soldeAvant := monCompte.GetSolde()
	test.Type(edtMontant, "100")
    test.Tap(btnCrediter)
    soldeApres := monCompte.GetSolde()
    if soldeApres - soldeAvant != float64(100) {
		t.Errorf("Solde incorrect : '%f', attendu : '%f'", soldeApres, soldeAvant+100)
	}
}

func TestDebiterCompte(t *testing.T) {
    CreerFenetreCompte()

    soldeAvant := monCompte.GetSolde()
	test.Type(edtMontant, "100")
    test.Tap(btnDebiter)
    soldeApres := monCompte.GetSolde()
    if soldeApres - soldeAvant != float64(-100) {
		t.Errorf("Solde incorrect : '%f', attendu : '%f'", soldeApres, soldeAvant-100)
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
