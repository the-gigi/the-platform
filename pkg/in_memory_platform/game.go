package in_memory_platform

import (
	"github.com/pkg/errors"
	om "github.com/the-gigi/the-platform/pkg/object_model"
)

type Game struct {
	om.Game
	state   string
	players map[string]bool
}

func (g Game) IsOpen() bool {
	return len(g.players) < g.Game.Type.MaxPlayerCount
}

func (g Game) Join(player string) (err error) {
	if !g.IsOpen() {
		err = errors.New("the game is not open")
		return
	}

	if g.players[player] {
		err = errors.New("player has already joined")
		return
	}

	g.players[player] = true
	g.PlayerCount++

	return
}

func newGame(gameId string, gameType om.GameType, state string) Game {
	return Game{
		Game: om.Game{
			Id:   gameId,
			Type: gameType,
		},
		players: map[string]bool{},
		state:   state,
	}
}
