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
		dominionGame      Game
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

		dominionGame = newGame("1", *dominionGameType, "")
		dominionGame.state = dominionGameState
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

	FContext("GameLobby tests", func() {

		It("should load state successfully", func() {
			_, err = p.CreateGame(*dominionGameType, "")
			Ω(err).Should(BeNil())

			state, err := p.LoadState(dominionGame.Id)
			Ω(err).Should(BeNil())
			Ω(state).Should(Equal(dominionGameState))
		})

		It("should save state successfully", func() {
		})

		It("should start game successfuly", func() {
		})

		It("should send data to a playersuccessfuly", func() {
		})

	})
})
