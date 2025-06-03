package main

import (
	"fmt"
	"variabler/gin_api_endpoint"
	variable "variabler/variableFolder"

	"github.com/gin-gonic/gin"
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

	//----API uden gin-----
	// routesAPI.MyRoutes()
	// http.ListenAndServe(":8080", nil)

	//API kald med gin
	ro := gin.Default()
	gin_api_endpoint.SetupRoutes(ro)
	ro.Run(":8080")

}
