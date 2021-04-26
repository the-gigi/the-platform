package in_memory_platform

import (
	"strconv"

	"github.com/pkg/errors"
	"github.com/the-gigi/the-platform/pkg/object_model/game_client"
	"github.com/the-gigi/the-platform/pkg/object_model/game_engine"
)

type InMemoryPlatform struct {
	Users        map[string]bool
	engines      map[game_engine.GameEngine]bool
	gameTypes    map[game_client.GameType]bool
	runningGames map[string]Game
	openGames    map[string]Game
	nextGameId   int
}

// GameLobby interface

func (p *InMemoryPlatform) GetGameTypes() (gameTypes []game_client.GameType, err error) {
	for gameType := range p.gameTypes {
		gameTypes = append(gameTypes, gameType)
	}
	return
}

func (p *InMemoryPlatform) GetOpenGames(gameType game_client.GameType) (games []game_client.Game, err error) {
	for _, game := range p.openGames {
		if game.Type == gameType {
			games = append(games, game.Game)
		}
	}
	return
}

func (p *InMemoryPlatform) CreateGame(gameType game_client.GameType) (gameId string, err error) {
	if !p.gameTypes[gameType] {
		err = errors.New("Unknown game type")
		return
	}

	p.nextGameId++
	gameId = strconv.Itoa(p.nextGameId)
	p.openGames[gameId] = newGame(gameId, gameType)

	return
}

func (p *InMemoryPlatform) Join(userId string, gameId string) (err error) {
	if !p.Users[userId] {
		err = errors.New("Unknown user")
		return
	}

	game, ok := p.openGames[gameId]
	if !ok {
		err = errors.Errorf("Game %s is not open", gameId)
		return
	}

	if game.PlayerCount == game.Type.MaxPlayerCount {
		err = errors.Errorf("Game %s is at capacity", gameId)
		return
	}

	game.players[userId] = true
	return
}

// GameEnginePlatform

func newInMemoryPlatform() *InMemoryPlatform {
	return &InMemoryPlatform{
		Users:        map[string]bool{},
		engines:      map[game_engine.GameEngine]bool{},
		gameTypes:    map[game_client.GameType]bool{},
		runningGames: map[string]Game{},
		openGames:    map[string]Game{},
	}

}
