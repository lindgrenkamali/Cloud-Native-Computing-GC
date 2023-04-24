package structs

import (
	"math/rand"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
	db, err := gorm.Open(sqlite.Open("db/gc"), &gorm.Config{})
	if err != nil {
		return false
	}

	db.AutoMigrate(&Team{})

	var count int64
	db.Model(&Team{}).Count(&count)

	if count < 1 {
		
	
	var cities []City

	db.Find(&cities)

	for i := 0; i < 20; i++ {

		citiesCount := len(cities)
		randCityId := cities[rand.Intn(citiesCount)].ID
		var city City
		db.First(&city, randCityId)

		randName := names[rand.Intn(len(names))]

		team := Team{Name: city.Name + " " + randName, CityID: city.ID, FoundedYear: rand.Intn(40) + 1980}

		db.Create(&team)
	}
	return true
}
return false
}

func RandomTeamID() int64 {
	db, err := gorm.Open(sqlite.Open("db/gc"), &gorm.Config{})
	if err != nil {

	}

	var teamCount int64

	db.Model(&Team{}).Count(&teamCount)

	return rand.Int63n(teamCount)
}

func ReturnAllTeams() []Team {
	db, err := gorm.Open(sqlite.Open("db/gc"), &gorm.Config{})
	if err != nil {
		return []Team{}
	}
	var teams []Team

	db.Preload("City").Find(&teams)
	return teams
}

func ReturnTeamByID(id string) Team {
	db, err := gorm.Open(sqlite.Open("db/gc"), &gorm.Config{})
	if err != nil {
		return Team{}
	}

	var team Team

	db.Preload("Players").Preload("City").First(&team, id)

	for i := 0; i < len(team.Players); i++ {
		team.Players[i].Team = nil
	}
	return team
}
