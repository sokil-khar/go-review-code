package business

import (
	"time"
	"fmt"
	"sort"
	"strings"
	"onefootball/model"
	"onefootball/modules"
)

var names = []string{"Germany", "England", "France", "Spain", "Manchester Utd", "Arsenal", "Chelsea", "Barcelona", "Real Madrid", "FC Bayern Munich"}

func GetTeam(id int) (model.Output, bool) {
	req := fmt.Sprintf("https://vintagemonster.onefootball.com/api/teams/en/%v.json", id)

	var output model.Output
	modules.GetRequest(req, &output)

	if name := output.Data.Team.Name; modules.Contains(names, name) {
		return output, true
	}

	return model.Output{}, false
}

func Print(output []model.Output){
	teamDetails := getTeamsFromOutput(output)
	aggregatePlayers(teamDetails)
}

func aggregatePlayers(teamDetails []model.Team){
	var players []model.Player

	for _, teamValue := range teamDetails {

		var p []model.Player
		for _, playerValue := range teamValue.Players {
			playerValue.TeamId = teamValue.Id
			playerValue.TeamName = teamValue.Name
			p = append(p, playerValue)
		}

		players = append(players, p...)
	}

	aggregate := make(map[string][]model.Player)

	for _, v := range players {
		aggregate[v.Id] = append(aggregate[v.Id], model.Player{v.Id, v.Name, v.FirstName, v.LastName, v.BirthDate, v.TeamId, v.TeamName})
	}

	print(aggregate)

}

func sortArray(players []finalResult) {
	sort.Slice(players, func(i, j int) bool {
		return players[i].firstName < players[j].firstName
	})
}

type finalResult struct {
	firstName string
	lastName string
	birthDate string
	teamMembers string
}

func print(aggregatePlayers map[string][]model.Player,){
	var finalResults []finalResult

	for k := range aggregatePlayers {
		p := aggregatePlayers[k][0]

		var teamNames []string
		for _, v := range aggregatePlayers[k] {
			teamNames = append(teamNames, v.TeamName)
		}

		finalResults = append(finalResults, finalResult{ p.FirstName, p.LastName, p.BirthDate, strings.Join(teamNames, ",")})
	}

	sortArray(finalResults)

	i := 1
	for _, v := range finalResults {
		fmt.Printf("%v. %v %v; %v; %v;\n", i, v.firstName, v.lastName, dateToAge(v.birthDate), v.teamMembers)
		i++
	}
}

func dateToAge(date string) int {
	t, _ := time.Parse("2006-01-02", date)
	now := time.Now()

	return now.Year() - t.Year()
}

func getTeamsFromOutput(output []model.Output) []model.Team {
	var teamDetails []model.Team

	for _,v := range output {
		teamDetails = append(teamDetails, v.Data.Team)
	}

	return teamDetails
}