package model

type Output struct {
	Code    int    `json:"code"`
	Data    Data   `json:"data"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type Data struct {
	Team Team `json:"team"`
}

type Team struct {
	Id         int      `json:"id"`
	Name       string   `json:"name"`
	IsNational bool     `json:"isNational"`
	Players    []Player `json:"players"`
}
type Player struct {
	Id        string    `json:"id"`
	Name      string `json:"name"`
	FirstName      string `json:"firstname"`
	LastName  string `json:"lastName"`
	BirthDate string `json:"birthDate"`
	TeamId int
	TeamName string
}
