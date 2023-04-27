package main

import (
	"testing"
	"io/ioutil"
	"os"
)

func TestAfficherCompte1(t *testing.T) {
    compte1 := CreerCompte("Bill Gates", 105300000000, "USD");
    attendu := "Bill Gates - 105300000000.000000 USD\n"

	rescueStdout := os.Stdout
  	r, w, _ := os.Pipe()
  	os.Stdout = w

  	compte1.Afficher()

  	w.Close()
  	obtenu, _ := ioutil.ReadAll(r)
  	os.Stdout = rescueStdout

    if string(obtenu) != attendu {
        t.Errorf("obtenu : '%s', attendu : '%s'", obtenu, attendu)
    }
}

func TestAfficherCompte2(t *testing.T) {
    compte2 := CreerCompte("Moi", 42, "EUR");
    attendu := "Moi - 42.000000 EUR\n"

	rescueStdout := os.Stdout
  	r, w, _ := os.Pipe()
  	os.Stdout = w

  	compte2.Afficher()

  	w.Close()
  	obtenu, _ := ioutil.ReadAll(r)
  	os.Stdout = rescueStdout

    if string(obtenu) != attendu {
        t.Errorf("obtenu : '%s', attendu : '%s'", obtenu, attendu)
    }
}

func TestAfficherCompte3(t *testing.T) {
    compte3 := CreerCompte("Maria", 1000, "YEN");
    attendu := "Maria - 1000.000000 YEN\n"

	rescueStdout := os.Stdout
  	r, w, _ := os.Pipe()
  	os.Stdout = w

  	compte3.Afficher()

  	w.Close()
  	obtenu, _ := ioutil.ReadAll(r)
  	os.Stdout = rescueStdout

    if string(obtenu) != attendu {
        t.Errorf("obtenu : '%s', attendu : '%s'", obtenu, attendu)
    }
}
