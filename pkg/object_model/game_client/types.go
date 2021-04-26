package game_client

type GameType struct {
	Name           string
	Description    string
	MinPlayerCount int
	MaxPlayerCount int
}

type Game struct {
	Id          string
	Type        GameType
	PlayerCount int
}
