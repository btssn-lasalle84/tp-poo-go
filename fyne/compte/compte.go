package main

import (
	"fmt"
	"image/color"

	//	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// Le concept Compte
// TODO

func FloatToString(valeur float64) string {
    return strconv.FormatFloat(valeur, 'f', 2, 64)
}

func StringToFloat(valeur string) float64 {
    conversion, _ := strconv.ParseFloat(valeur, 64)
    return conversion
}

// Le compte
var monCompte Compte

// Les widgets
var txtTitre *canvas.Text
var txtSolde *canvas.Text
var txtMontant *canvas.Text
var edtMontant *widget.Entry
var btnCrediter *widget.Button
var btnDebiter *widget.Button

func CreerFenetreCompte() *fyne.Container {
    monCompte = CreerCompte("Moi", 42, "EUR")

    // TODO

	//return container.New(layout.NewVBoxLayout(), ...)
}

func main() {
    applicationCompte := app.New()
	fenetreCompte := applicationCompte.NewWindow("Compte")
	fenetreCompte.SetContent(CreerFenetreCompte())
	fenetreCompte.ShowAndRun()
}
