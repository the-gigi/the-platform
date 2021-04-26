package in_memory_platform

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/the-gigi/the-platform/pkg/mocks"
	om "github.com/the-gigi/the-platform/pkg/object_model"
)

const (
	dominionGameState = "dominion-state-1"
)

var _ = Describe("Platform Tests", func() {
	var (
		p                 *InMemoryPlatform
		err               error
		gameTypes         []om.GameType
		dominionGameType  *om.GameType
		blocktserGameType *om.GameType

		dominionGame      *Game
		mockGameEngine    *mocks.MockGameEngine
	)

	BeforeEach(func() {
		p = newInMemoryPlatform()
		err = nil
		gameTypes = []om.GameType{}
		dominionGameType = &om.GameType{
			Name:           "dominion",
			Description:    "card game",
			MinPlayerCount: 2,
			MaxPlayerCount: 6,
		}
		blocktserGameType = &om.GameType{
			Name:           "blocktser",
			Description:    "block puzzle",
			MinPlayerCount: 1,
			MaxPlayerCount: 1,
		}

		dominionGame, err = newGame("1", *dominionGameType, dominionGameState)
		Ω(err).Should(BeNil())
		mockGameEngine = &mocks.MockGameEngine{}
	})

	Context("GameLobby tests", func() {
		It("should return game types correctly", func() {
			gameTypes, err = p.GetGameTypes()
			Ω(err).Should(BeNil())
			Ω(gameTypes).Should(HaveLen(0))

			p.gameTypes[dominionGameType.Name] = dominionGameType
			p.gameTypes[blocktserGameType.Name] = blocktserGameType

			gameTypes, err = p.GetGameTypes()
			Ω(err).Should(BeNil())
			Ω(gameTypes).Should(HaveLen(2))
		})

		It("should register game successfully", func() {
			err = p.Register(*dominionGameType, mockGameEngine)
			Ω(err).Should(BeNil())

			Ω(p.engines[mockGameEngine]).Should(Equal(dominionGameType))
		})

		It("should create game successfully", func() {
			err = p.Register(*dominionGameType, mockGameEngine)
			Ω(err).Should(BeNil())

			_, err = p.CreateGame(*dominionGameType, dominionGame.state)
			Ω(err).Should(BeNil())

			Ω(p.openGames).Should(HaveLen(1))
			Ω(p.openGames[dominionGame.Id]).Should(Equal(dominionGame))
		})
	})

	Context("Platform tests", func() {

		It("should load state successfully", func() {
			err = p.Register(*dominionGameType, mockGameEngine)
			Ω(err).Should(BeNil())

			p.runningGames[dominionGame.Id] = dominionGame

			state, err := p.LoadState(dominionGame.Id)
			Ω(err).Should(BeNil())
			Ω(state).Should(Equal(dominionGameState))

			newState := "new-state"
			gameId, err := p.CreateGame(*dominionGameType, newState)
			Ω(err).Should(BeNil())

			err = p.startGame(gameId)
			Ω(err).Should(BeNil())

			state, err = p.LoadState(gameId)
			Ω(err).Should(BeNil())
			Ω(state).Should(Equal(newState))
		})

		It("should save state successfully", func() {
			p.runningGames[dominionGame.Id] = dominionGame

			state, err := p.LoadState(dominionGame.Id)
			Ω(err).Should(BeNil())
			Ω(state).Should(Equal(dominionGameState))

			newState := "new-state"
			err = p.SaveState(dominionGame.Id, newState)
			Ω(err).Should(BeNil())

			state, err = p.LoadState(dominionGame.Id)
			Ω(err).Should(BeNil())
			Ω(state).Should(Equal(newState))
		})

		It("should start game successfully", func() {
		})

		It("should send data to a player successfully", func() {
		})
	})
})
