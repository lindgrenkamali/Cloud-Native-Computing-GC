package main

import (
	persons "GC/apis"
	ginning "GC/gin"
	"GC/structs"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

var Global = 20

func main() {
	
	r := gin.Default()
	structs.Config = GetConfig()
	structs.AddCities()
	structs.AddRandomTeams()
	structs.AddRandomPlayersForTeams()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", ginning.Home)
	r.GET("/api/team", ginning.Team)
	r.GET("/api/david", persons.David)
	r.GET("/api/team/:id", ginning.TeamById)
	r.GET("/api/player/:id", ginning.PlayerById)
	r.Run(":8080")

}

func GetConfig()structs.ConfigFile{
	fileName := "Config.yml"

	e := os.Getenv("RUNENVIRONMENT")

	if len(e) > 0 {
	
		fileName = "production" + fileName

	} else {
		fileName = "dev" + fileName
	}

	f, _ := os.Open(fileName)
	defer f.Close()
	decoder := yaml.NewDecoder(f)
	var config structs.ConfigFile
	decoder.Decode(&config)
	return config
}