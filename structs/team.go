package structs

import (
	"GC/dbcontext"
	"math/rand"
)

type Team struct {
	ID          int
	Name        string
	City        City
	CityID      int
	FoundedYear int
	Players     []Player
}

var names = [20]string{"Hawks", "Eagels", "Bulls", "Titans",
	"Sharks", "Goats", "Vikings", "Lions", "Giants", "Falcons",
	"Parrots", "Dolphins", "Panthers", "Jaguars", "Capybaras",
	"Dinosaurs", "Rocks", "Meteors", "Alligators", "Raiders"}

func AddRandomTeams() bool {


	dbcontext.DB.AutoMigrate(&Team{})

	var count int64
	dbcontext.DB.Model(&Team{}).Count(&count)

	if count < 1 {
		
	
	var cities []City

	dbcontext.DB.Find(&cities)

	for i := 0; i < 20; i++ {

		citiesCount := len(cities)
		randCityId := cities[rand.Intn(citiesCount)].ID
		var city City
		dbcontext.DB.First(&city, randCityId)

		randName := names[rand.Intn(len(names))]

		team := Team{Name: city.Name + " " + randName, CityID: city.ID, FoundedYear: rand.Intn(40) + 1980}

		dbcontext.DB.Create(&team)
	}
	

	return true
}
return false
}


func RandomTeamID() int64 {

	var teamCount int64

	dbcontext.DB.Model(&Team{}).Count(&teamCount)

	return rand.Int63n(teamCount)
}

func ReturnAllTeams() []Team {

	var teams []Team

	dbcontext.DB.Preload("City").Find(&teams)
	return teams
}

func ReturnTeamByID(id string) Team {

	var team Team

	dbcontext.DB.Preload("Players").Preload("City").First(&team, id)

	for i := 0; i < len(team.Players); i++ {
		team.Players[i].Team = nil
	}
	return team
}
