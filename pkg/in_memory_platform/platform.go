package in_memory_platform

import (
	"strconv"

	"github.com/pkg/errors"
	"github.com/the-gigi/the-platform/pkg/object_model/game_engine"

	om "github.com/the-gigi/the-platform/pkg/object_model"
)

type GameType = string

type InMemoryPlatform struct {
	users         map[string]bool
	engines       map[game_engine.GameEngine]*om.GameType
	gameTypes     map[string]*om.GameType
	openGames     map[string]Game
	runningGames  map[string]Game
	finishedGames map[string]Game

	nextGameId int
}

// GameLobby interface
func (p *InMemoryPlatform) GetGameTypes() (gameTypes []om.GameType, err error) {
	for _, gameType := range p.gameTypes {
		gameTypes = append(gameTypes, *gameType)
	}
	return
}

func (p *InMemoryPlatform) GetOpenGames(gameType om.GameType) (games []om.Game, err error) {
	for _, game := range p.openGames {
		if game.Type == gameType {
			games = append(games, game.Game)
		}
	}
	return
}

func (p *InMemoryPlatform) CreateGame(gameType om.GameType, state string) (gameId string, err error) {
	if p.gameTypes[gameType.Name] == nil {
		err = errors.Errorf("Unknown game type: %s", gameType.Name)
		return
	}

	p.nextGameId++
	gameId = strconv.Itoa(p.nextGameId)
	p.openGames[gameId] = newGame(gameId, gameType, state)

	return
}

// Platform interface
func (p *InMemoryPlatform) isGameTypeValid(gameType om.GameType) (err error) {
	if gameType.Name == "" {
		err = errors.New("game type name can't be empty")
		return
	}

	if gameType.Description == "" {
		err = errors.New("game type description can't be empty")
		return
	}

	if gameType.MinPlayerCount <= 0 {
		err = errors.New("min player count must be at least one")
		return
	}

	if gameType.MinPlayerCount > gameType.MaxPlayerCount {
		err = errors.New("min player count must be less than or equal to max player count")
		return
	}

	if p.gameTypes[gameType.Name] != nil {
		err = errors.Errorf("there is already a registered game type named: %s", gameType.Name)
		return
	}
	return
}

func (p *InMemoryPlatform) Register(gameType om.GameType, gameEngine game_engine.GameEngine) (err error) {
	if p.engines[gameEngine] != nil {
		err = errors.Errorf("engine %v already registered for game type %s", gameEngine, gameType.Name)
		return
	}

	if p.gameTypes[gameType.Name] != nil {
		err = errors.Errorf("gameType %s is already registered", gameType.Name)
		return
	}

	err = p.isGameTypeValid(gameType)
	if err != nil {
		return
	}

	p.engines[gameEngine] = &gameType
	p.gameTypes[gameType.Name] = &gameType
	return
}

func (p *InMemoryPlatform) LoadState(gameId string) (state string, err error) {
	game, ok := p.runningGames[gameId]
	if !ok {
		err = errors.Errorf("no such game: %s", gameId)
		return
	}

	state = game.state
	return
}

func (p *InMemoryPlatform) SaveState(gameId string, state string) (err error) {
	game, ok := p.runningGames[gameId]
	if !ok {
		err = errors.Errorf("no such game: %s", gameId)
		return
	}

	game.state = state
	return
}

func (p *InMemoryPlatform) GameOver(gameId string) (err error) {
	game, ok := p.runningGames[gameId]
	if !ok {
		err = errors.Errorf("no such running game: %s", gameId)
		return
	}

	delete(p.runningGames, gameId)
	p.finishedGames[gameId] = game

	return
}

func (p *InMemoryPlatform) Send(gameEngine, player string, data string) (err error) {
	return
}

// GameEngine interface
func (p *InMemoryPlatform) Join(userId string, gameId string) (err error) {
	if !p.users[userId] {
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
		users:         map[string]bool{},
		engines:       map[game_engine.GameEngine]*om.GameType{},
		gameTypes:     map[string]*om.GameType{},
		openGames:     map[string]Game{},
		runningGames:  map[string]Game{},
		finishedGames: map[string]Game{},
	}
}
