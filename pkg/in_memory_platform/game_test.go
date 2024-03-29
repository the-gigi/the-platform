package in_memory_platform

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	om "github.com/the-gigi/the-platform/pkg/object_model"
)

var _ = Describe("Game Tests", func() {
	var (
		err  error
		game *Game
	)

	BeforeEach(func() {
		gameType := om.GameType{
			Name:           "cool game",
			Description:    "a cool game",
			MinPlayerCount: 2,
			MaxPlayerCount: 4,
		}
		game, err = newGame("1", gameType, "")
		Ω(err).Should(BeNil())
		game.PlayerCount = len(game.players)
	})

	Context("Join() tests", func() {
		It("should join successfully an open game", func() {
			Ω(game.Join("player-1")).Should(Succeed())
		})
		It("should fail when game is not open", func() {
			Ω(game.Join("player-1")).Should(Succeed())
			Ω(game.Join("player-2")).Should(Succeed())
			Ω(game.Join("player-3")).Should(Succeed())
			Ω(game.Join("player-4")).Should(Succeed())
			Ω(game.Join("player-5")).ShouldNot(Succeed())
		})
		It("should fail when the player has already joined", func() {
			Ω(game.Join("player-1")).Should(Succeed())
			Ω(game.Join("player-1")).ShouldNot(Succeed())
		})
	})

	Context("IsOpen() tests", func() {
		It("should return true when game is open", func() {
			Ω(game.IsOpen()).Should(Equal(true))
		})
		It("should return false when game is not open", func() {
			Ω(game.Join("player-1")).Should(Succeed())
			Ω(game.Join("player-2")).Should(Succeed())
			Ω(game.Join("player-3")).Should(Succeed())
			Ω(game.Join("player-4")).Should(Succeed())
			Ω(game.IsOpen()).Should(Equal(false))
		})
	})
})
