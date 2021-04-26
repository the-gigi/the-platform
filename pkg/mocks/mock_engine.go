package mocks

import (
	. "github.com/onsi/gomega"
	"github.com/the-gigi/sham"
)

// MockImporter a mock object for the importer
type MockGameEngine struct {
	sham.CannedResponseMock
	call sham.Call
}

func (g *MockGameEngine) OnGameStart(gameId string) {
	var err error
	g.call, err = g.VerifyCall("OnGameStart", 0, gameId)
	Ω(err).Should(BeNil())
	return
}

func (g *MockGameEngine) OnJoin(gameId string, playerId string) {
	var err error
	g.call, err = g.VerifyCall("OnGameStart", 0, gameId)
	Ω(err).Should(BeNil())
	return
}

func (g *MockGameEngine) OnLeave(gameId string, playerId string) {
	var err error
	g.call, err = g.VerifyCall("OnGameStart", 0, gameId)
	Ω(err).Should(BeNil())
	return
}

func (g *MockGameEngine) Do(gameId string, playerId string, action string, data string) (result string, err error) {
	g.call, err = g.VerifyCall("Do", 2, gameId, playerId, action, data)
	Ω(err).Should(BeNil())

	result = g.call.Result[0].(string)
	err = sham.ToError(g.call.Result[1])

	return
}
