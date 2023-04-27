package main

import "fmt"

func main() {
  var a int = 1  // un entier (ou int64)
  var b float64 = 2.5 // un nombre à virgule flottante (ou float)
  var c string = "hello" // une chaîne de caractères
  var d = false // un booléen
  e := "world" // déclare et initialise une variable
  const NB int = 10 // une constante
  var t[NB] int // un tableau d'entiers initialisés à 0
  var fruits = []string {"pomme", "banane", "fraise", "cerise"} // slice

  // afficher une variable :
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
  fmt.Println(c + " " + e)
  fmt.Println(NB)
  fmt.Println(t[0])
  fmt.Println(fruits)
  fmt.Println(len(fruits))

  // cast
  var i int = 42
  var f float64 = float64(i)
  var u uint = uint(f)
  fmt.Printf("i : %T\n", i)
  fmt.Printf("f : %T\n", f)
  fmt.Printf("u : %T\n", u)
}
