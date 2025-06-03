package gin_api_endpoint

import (
	"variabler/gin_api_funktioner"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(url *gin.Engine) {
	url.POST("/beregn-rente", gin_api_funktioner.RenteHandler)
}
