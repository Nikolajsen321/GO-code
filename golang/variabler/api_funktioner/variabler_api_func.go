package apifunktioner

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func RenteHandler(rs http.ResponseWriter, rq *http.Request) {
	query := rq.URL.Query()
	startStr := query.Get("start")
	renteStr := query.Get("rente")
	aarStr := query.Get("aar")

	start, _ := strconv.ParseFloat(startStr, 64)
	rente, _ := strconv.ParseFloat(renteStr, 64)
	aar, _ := strconv.Atoi(aarStr)

	slut := start
	for i := 0; i < aar; i++ {
		slut *= (1 + rente)
	}

	rs.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rs).Encode(map[string]float64{
		"slutkapital": slut,
	})

}
