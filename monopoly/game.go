package monopoly

import (
	"log"
)

type Game struct {
	remainingPlayers int
	turnQueue        TurnQueue
	turnCount        int
}

func InitGame(playerCount int) Game {
	var g Game

	g.turnCount = 1
	g.remainingPlayers = playerCount
	firstPlayer := Player{
		Id:      1,
		balance: 1500,
		square:  0,
	}

	firstPlayerTurn := &PlayerTurn{
		player: &firstPlayer,
		next:   nil,
	}

	g.turnQueue = TurnQueue{
		head: firstPlayerTurn,
		tail: firstPlayerTurn,
	}

	for i := 1; i < playerCount; i++ {
		player := Player{
			Id:      uint8(i) + 1,
			balance: 1500,
			square:  0,
		}

		g.turnQueue.Push(&player)
	}

	return g
}

func (g *Game) IsOver() *Player {
	if g.remainingPlayers == 1 {
		return g.turnQueue.Pop()
	} else {
		return nil
	}
}

func (g *Game) NextTurn() {
	player := g.turnQueue.Pop()
	
	log.Println()
	log.Printf("Turn #%d - P%d", g.turnCount, player.Id)
	
	canContinue := player.PlayTurn()
	if canContinue {
		g.turnQueue.Push(player)
	} else {
		log.Printf("P%d stopped on Go: they've been kicked out of the game!", player.Id)
		g.remainingPlayers--
	}

	g.turnCount++
}
