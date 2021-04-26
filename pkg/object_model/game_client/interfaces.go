package game_client

import (
	om "github.com/the-gigi/the-platform/pkg/object_model"
)

type GameLobby interface {
	GetGameTypes() (gameTypes []om.GameType, err error)
	GetOpenGames(gameType om.GameType) (games []om.Game, err error)
	CreateGame(gameType om.GameType, state string) (gameId string, err error)
	Join(gameId string) (playerId, err error)
	StartGame(playerId, gameId string) (err error)
}

type RunningGame interface {
	Do(playerId string, action string, data string) (result string, err error)
	Leave() (err error)
}

type GameClient interface {
	Receive(data string) error
}
