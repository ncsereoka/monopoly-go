package main

import (
	"flag"
	"log"
	"math/rand"
	"time"

	"monopoly/monopoly"
)

func main() {
	var playerCount int
	flag.IntVar(&playerCount, "players", 4, "The number of players: 2 to 6. Default is 4")
	flag.Parse()

	if playerCount < 2 || playerCount > 6 {
		log.Fatal("The number of players must be between 2 and 6!")
		return
	}

	rand.Seed(time.Now().UnixNano())
	var game monopoly.Game
	game.Init(playerCount)

	log.Printf("Starting a new game with %d players...\n", playerCount)

	var winner *monopoly.Player
	for ; winner == nil; winner = game.IsOver() {
		game.Next()
	}

	log.Println()
	log.Printf("P%d is the last player remaining. P%d has won the game!", winner.Id, winner.Id)
}
