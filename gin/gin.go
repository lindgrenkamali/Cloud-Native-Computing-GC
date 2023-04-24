package ginning

import (
	"GC/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", nil)
}

func Team(c *gin.Context) {
	c.JSON(http.StatusOK, structs.ReturnAllTeams())
}

func TeamById(c *gin.Context) {
	stringId := c.Param("id")

	c.JSON(http.StatusOK, structs.ReturnTeamByID(stringId))
}

func PlayerById(c *gin.Context) {
	stringId := c.Param("id")

	c.JSON(http.StatusOK, structs.ReturnPlayerByID(stringId))
}
