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
	PlayerMap[1] = &firstPlayer

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
			Id:      i + 1,
			balance: 1500,
			square:  0,
		}
		PlayerMap[i+1] = &player

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

	// log.Println()
	// log.Printf("Turn #%d - P%d on square#%d", g.turnCount, player.Id, player.square)

	canContinue := player.PlayTurn()
	if canContinue {
		g.turnQueue.Push(player)
	} else {
		log.Printf("P%d has more than $10000: they've been kicked out of the game!", player.Id)
		g.remainingPlayers--
	}

	g.turnCount++
}
