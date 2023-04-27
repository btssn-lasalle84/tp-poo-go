package main

import (
	"testing"
	"io/ioutil"
	"os"
)

func TestNomPoint(t *testing.T) {
	p1 := Point{"p1", 2, 5}
    attendu := "p1"
    obtenu := p1.Nom
    if obtenu != attendu {
        t.Errorf("obtenu : '%s', attendu : '%s'", obtenu, attendu)
    }
}

func TestXPoint(t *testing.T) {
	p1 := Point{"p1", 2, 5}
    attendu := 2
    obtenu := p1.X
    if obtenu != attendu {
        t.Errorf("obtenu : '%d', attendu : '%d'", obtenu, attendu)
    }
}

func TestYPoint(t *testing.T) {
	p1 := Point{"p1", 2, 5}
    attendu := 5
    obtenu := p1.Y
    if obtenu != attendu {
        t.Errorf("obtenu : '%d', attendu : '%d'", obtenu, attendu)
    }
}

func TestAffichePoint(t *testing.T) {
	p1 := Point{"p1", 2, 5}
    attendu := "p1(2,5)\n"

	rescueStdout := os.Stdout
  	r, w, _ := os.Pipe()
  	os.Stdout = w

  	p1.Affiche()

  	w.Close()
  	obtenu, _ := ioutil.ReadAll(r)
  	os.Stdout = rescueStdout

    if string(obtenu) != attendu {
        t.Errorf("obtenu : '%s', attendu : '%s'", obtenu, attendu)
    }
}

func TestDistancePoint(t *testing.T) {
	p1 := Point{"p1", 2, 5}
    p2 := CreerPoint()
    attendu := 5.385164807134504
    obtenu := p1.Distance(p2);
    if obtenu != attendu {
        t.Errorf("obtenu : '%f', attendu : '%f'", obtenu, attendu)
    }
	t.Log("La distance entre deux points")
}

func TestCoordonneesPoint(t *testing.T) {
	p1 := Point{"p1", 2, 5}
	p1.SetCoordonnees(3, 6)
    attendu := 3
    obtenu := p1.X
    if obtenu != attendu {
        t.Errorf("obtenu : '%d', attendu : '%d'", obtenu, attendu)
    }
	attendu = 6
    obtenu = p1.Y
    if obtenu != attendu {
        t.Errorf("obtenu : '%d', attendu : '%d'", obtenu, attendu)
    }
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
