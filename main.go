package main

import (
	persons "GC/apis"
	ginning "GC/gin"
	"GC/structs"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	structs.AddCities()
	structs.AddRandomTeams()
	structs.AddRandomPlayersForTeams()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", ginning.Home)
	r.GET("/api/david", persons.David)
	r.Run(":8080")


}