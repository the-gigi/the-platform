package game_client

type GameLobby interface {
	GetGameTypes() (gameTypes []GameType, err error)
	GetOpenGames(gameType GameType) (games []Game, err error)
	CreateGame(gameType GameType) (gameId string, err error)
	Join(gameId string) (err error)
}

type RunningGame interface {
	Do(player string, action string, data string) (result string, err error)
	Leave() (err error)
}

type GameClient interface {
	Receive(data string) error
}
