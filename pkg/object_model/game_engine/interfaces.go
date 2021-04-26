package game_engine

type Platform interface {
	Register(gameType string, gameEngine GameEngine) error
	LoadState(stateId string) (state string, err error)
	SaveState(state string) (stateId string, err error)
	StartGame(gameEngine GameEngine) (gameId string, err error)
	GameOver(gameEngine GameEngine) error
	Send(gameEngine, player string, data string) error
}

type GameEngine interface {
	Join(player string) error
	Leave(player string) error
	Do(player string, action string, data string) (result string, err error)
}
