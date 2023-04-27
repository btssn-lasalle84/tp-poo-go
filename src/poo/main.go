package main

import (
  "fmt"
  "lasalle.com/poo/point"
)

// Une fonction acceptant un argument de l'interface IPoint
func foo(p point.IPoint) {
  fmt.Println(p)
	p.Affiche()
}

func main() {
	p1 := point.Point{"p1", 2, 5}
  fmt.Println("X =", p1.X)
  fmt.Println("Y =", p1.Y)
	p1.Affiche()
  p2 := point.CreerPoint()
  p2.Affiche()
  distance := p1.Distance(p2);
  fmt.Println("distance =", distance)
  p1.SetCoordonnees(3, 6)
  foo(p1)
}
