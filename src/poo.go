package main

import (
  "fmt"
	"math"
)

// Une interface
// Ensemble de méthodes sur le type Point
type IPoint interface {
  Affiche()
  Distance(p Point) float64
  SetCoordonnees(x int, y int)
}

// Une structure Point
type Point struct {
  Nom string
  X   int
  Y   int
}

// Définit la méthode Affiche() liée à la structure Point
func (point Point) Affiche() {
  fmt.Printf("%s(%d,%d)\n", point.Nom, point.X, point.Y)
}

// Une méthode Distance()
func (point Point) Distance(p Point) float64 {
  var dx float64 = float64(point.X - p.X)
  var dy float64 = float64(point.Y - p.Y)
  return math.Sqrt(dx * dx + dy * dy)
}

// Définit la méthode SetCoordonnees() liée à la structure Point
// ici on utilise un pointeur pour modifier la variable point
func (point* Point) SetCoordonnees(x int, y int) {
  point.X = x
  point.Y = y
}

// Une fonction
// Crée et retourne une instance par défaut de Point
func CreerPoint() Point {
  point := Point{"", 0, 0}
  return point
}

// Une fonction acceptant un argument de l'interface IPoint
func foo(point IPoint) {
  fmt.Println(point)
	point.Affiche()
}

func main() {
	p1 := Point{"p1", 2, 5}
  fmt.Println("X =", p1.X)
  fmt.Println("Y =", p1.Y)
	p1.Affiche()
  p2 := CreerPoint()
  p2.Affiche()
  distance := p1.Distance(p2);
  fmt.Println("distance =", distance)
  p1.SetCoordonnees(3, 6)
  foo(p1)
}
