package game_engine

import (
	om "github.com/the-gigi/the-platform/pkg/object_model"
)

type Platform interface {
	Register(gameType om.GameType, gameEngine GameEngine) error
	LoadState(gameId string) (state string, err error)
	SaveState(gameId string, state string) (err error)
	GameOver(gameId string) error
	Send(gameId string, player string, data string) error
}

type GameEngine interface {
	OnGameStart(gameId string)
	OnJoin(gameId string, playerId string)
	OnLeave(gameId string, playerId string)
	Do(gameId string, playerId string, action string, data string) (result string, err error)
}
