package main

import (
	"fmt"
	variable "variabler/variableFolder"
)

func main() {
	r := variable.Rectangle{Height: 10, Width: 5}

	fmt.Println("Areal ", r.Area())
	area := variable.RectangleArea(r)
	fmt.Println("Areal:", area)
	startKapital := 1000.0
	rente := 0.5
	aar := 3

	var slutKapital float64 = variable.RenteRegning(startKapital, rente, aar)
	var slutKapital2 float64 = variable.RenteRegningRecur(startKapital, rente, aar)

	fmt.Printf("Hvis du indsætter %v med en rente på %v over %v år  får man %v",
		startKapital, rente, aar, slutKapital)

	fmt.Println("\nRekursive:")

	fmt.Printf("Hvis du indsætter %v med en rente på %v over %v år  får man %v",
		startKapital, rente, aar, slutKapital2)

}

/*
GO comandos:
go run main.go  fx hvis man er i mappen
go run ../main.go  fx hvis man er i en sub mappe af den mappe main ligger i
go run ./  hvis mappen indholder en main package
*/
