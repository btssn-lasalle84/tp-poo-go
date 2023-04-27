package main

import "fmt"

// Le programme
func main() {
  var langue string = "fr"
  var message string
  fmt.Print("Quelle est votre langue (fr, ...) ? ");
  // saisie d'une chaîne de caractères
  fmt.Scanf("%s", &langue)

  // une instruction conditionnelle
  if (langue == "fr") {
    message = "Bonjour le monde";
  } else {
    message = "Hello world";
  }

  // saisie d'un entier
  var nb int = 5
  fmt.Print("Donnez un nombre : ");
  fmt.Scanf("%d", &nb)

  var i int
  // une boucle
  for i = 0; i < nb; i++ {
    fmt.Printf("%s %d fois\n", message, (i+1))
  }
}
