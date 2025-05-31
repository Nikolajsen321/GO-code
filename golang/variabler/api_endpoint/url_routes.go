package routesAPI

import (
	"net/http"
	apifunktioner "variabler/api_funktioner"
)

func MyRoutes() {
	http.HandleFunc("/beregn-rente", apifunktioner.RenteHandler)

}
