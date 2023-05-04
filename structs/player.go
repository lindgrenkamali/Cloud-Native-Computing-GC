package structs

import (
	"GC/dbcontext"
	"math/rand"
)

type Player struct {
	ID           int
	Name         string
	Team         *Team
	TeamID       int
	JerseyNumber int
	BirthYear    int
}

var forenames = [20]string{"Anna", "Lisa", "Emil", "Carl", "Rick",
	"Elon", "Steve", "Juli", "Filip", "Stefan", "Anders", "Annika",
	"Lars", "Mick", "Carola", "Thea", "Lola", "Stig", "Josefin", "Helmer"}

var surnames = [20]string{"Lindgren", "Lindqvist", "Larsson", "Eliasson", "Häggkvist", "Andersson", "Svensson",
	"Höglund", "Löfven", "Thunberg", "Gustafsson", "Djarin", "Skywalker", "Jones", "Wayne", "Ågren", "Munther",
	"Maggio", "Timell", "Johansson"}

func randomName() (fullname string) {
	randomNumber1 := rand.Int63n(20)
	randomNumber2 := rand.Int63n(20)

	forename := forenames[randomNumber1]
	surname := surnames[randomNumber2]

	fullname = forename + " " + surname

	return
}

func AddRandomPlayerWithTeamID(id int) bool {

	dbcontext.DB.AutoMigrate(&Player{})

	fullName := randomName()
	teamId := id

	jerseyNumber := rand.Int63n(100)
	birthYear := rand.Int63n(5) + 2000

	p := Player{Name: fullName, TeamID: int(teamId), JerseyNumber: int(jerseyNumber), BirthYear: int(birthYear)}

	dbcontext.DB.Create(&p)

	return true
}

func AddRandomPlayersForTeams() bool {

	dbcontext.DB.AutoMigrate(&Player{})

	var count int64
	dbcontext.DB.Model(&Player{}).Count(&count)

	if count < 1 {
		
	var teams []Team

	dbcontext.DB.Find(&teams)

	for i := 0; i < len(teams); i++ {
		for j := 0; j < 5; j++ {
			AddRandomPlayerWithTeamID(teams[i].ID)
		}

	}

	dbcontext.DB.AutoMigrate(&Player{})

	return true
}
return false
}

func ReturnPlayerByID(id string) Player {

	var player Player

	dbcontext.DB.Preload("Team").First(&player, id)

	return player
	
}
