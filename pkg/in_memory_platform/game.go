package in_memory_platform

import (
	"github.com/pkg/errors"
	"github.com/the-gigi/the-platform/pkg/object_model/game_client"
)

type Game struct {
	game_client.Game
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

func newGame(gameId string, gameType game_client.GameType) Game {
	return Game{
		Game: game_client.Game{
			Id:   gameId,
			Type: gameType,
		},
		players: map[string]bool{},
	}
}
