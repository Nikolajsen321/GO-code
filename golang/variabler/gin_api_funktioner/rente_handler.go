package gin_api_funktioner

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RenteInput struct {
	Start float64
	Rente float64
	Aar   int
}

type RenteResult struct {
	Slutkapital float64 `json:"slutkapital"`
}

func RenteHandler(c *gin.Context) {
	var input RenteInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ugyldigt input"})
		return
	}
	slut := input.Start
	for i := 0; i < input.Aar; i++ {
		slut *= (1 + input.Rente)
	}
	c.JSON(http.StatusOK, RenteResult{Slutkapital: slut})

}
