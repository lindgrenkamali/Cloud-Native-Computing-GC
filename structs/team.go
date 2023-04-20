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
}

var names = [20]string{"Hawks", "Eagels", "Bulls", "Titans",
	"Sharks", "Goats", "Vikings", "Lions", "Giants", "Falcons",
	"Parrots", "Dolphins", "Panthers", "Jaguars", "Capybaras",
	"Dinosaurs", "Rocks", "Meteors", "Alligators", "Raiders"}

func AddRandomTeam()bool {
	db, err := gorm.Open(sqlite.Open("db/gc.db"), &gorm.Config{})
	if err != nil {
		return false
	}

	db.AutoMigrate(&Team{})
	var count int64


	db.Model(&City{}).Count(&count)

	randCityId := rand.Int63n(count)
	var city City
	db.First(&city, randCityId)

	randName := names[rand.Intn(len(names))]

	team := Team {Name: city.Name + " " + randName, CityID: city.Id, FoundedYear: rand.Intn(40) + 1980}

	db.Create(&team)

	return true
}

func RandomTeamID() int64{
	db, err := gorm.Open(sqlite.Open("db/gc.db"), &gorm.Config{})
	if err != nil {
		
	}

	var teamCount int64

	db.Model(&Team{}).Count(&teamCount)

	return rand.Int63n(teamCount)
}