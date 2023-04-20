package persons

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string
	City string
} 

func David(c *gin.Context) {
	david := Person {"David", "Halmstad"}

	c.JSON(http.StatusOK, david)
}