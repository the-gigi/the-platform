package in_memory_platform

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/the-gigi/the-platform/pkg/object_model/game_client"
)

var _ = Describe("Platform Tests", func() {
	var (
		p                 *InMemoryPlatform
		err               error
		gameTypes         []game_client.GameType
		dominionGameType  game_client.GameType
		blocktserGameType game_client.GameType
	)

	BeforeEach(func() {
		p = newInMemoryPlatform()
		err = nil
		gameTypes = []game_client.GameType{}
		dominionGameType = game_client.GameType{
			Name:           "dominion",
			Description:    "card game",
			MinPlayerCount: 2,
			MaxPlayerCount: 6,
		}
		blocktserGameType = game_client.GameType{
			Name:           "blocktser",
			Description:    "block puzzle",
			MinPlayerCount: 1,
			MaxPlayerCount: 1,
		}
	})

	Context("GameLobby tests", func() {
		It("should return game types correctly", func() {
			gameTypes, err = p.GetGameTypes()
			Ω(err).Should(BeNil())
			Ω(gameTypes).Should(HaveLen(0))

			p.gameTypes[dominionGameType] = true
			p.gameTypes[blocktserGameType] = true

			gameTypes, err = p.GetGameTypes()
			Ω(err).Should(BeNil())
			Ω(gameTypes).Should(HaveLen(2))
		})
	})
})
