package monopoly

import (
	"fmt"
	"log"
	"math/rand"
)

type Player struct {
	Id      uint8
	balance uint16
	square  uint8
	jailed  bool
	keys    uint8
}

func (p Player) String() string {
	return fmt.Sprintf("P%d with a balance of $%d is on square#%d ", p.Id, p.balance, p.square)
}

func singleDiceThrow() int {
	return rand.Intn(6) + 1
}

func tripleDiceThrow() int {
	return singleDiceThrow() + singleDiceThrow() + singleDiceThrow()
}

func (p *Player) takeAChance() {
	roll := tripleDiceThrow()
	firstPart := fmt.Sprintf("P%d takes a Chance and rolls %d:", p.Id, roll)

	switch roll {
	case 3:
		p.balance -= 15
		log.Printf("%s Pays a $15 speeding fine", firstPart)
		return
	case 4:
		p.balance += 50
		log.Printf("%s Collects $50 of bank dividends", firstPart)
		return
	case 5:
		log.Printf("%s Advances to Trafalgar Square", firstPart)
		return
	case 6:
		log.Printf("%s Advances to Go", firstPart)
		return
	case 7:
		p.balance -= 150
		log.Printf("%s Pays $150 of school fees", firstPart)
		return
	case 8:
		log.Printf("%s Advances to Mayfair", firstPart)
		return
	case 9:
		p.keys++
		log.Printf("%s Receives a Jail key", firstPart)
		return
	case 10:
		log.Printf("%s Is assessed for street repairs", firstPart)
		return
	case 11:
		log.Printf("%s Makes general repairs on all of their buildings", firstPart)
		return
	case 12:
		log.Printf("%s Advances to Pall Mall", firstPart)
		return
	case 13:
		log.Printf("%s Goes to Marylebone Station", firstPart)
		return
	case 14:
		p.balance += 150
		log.Printf("%s Collects $150 from a matured building loan", firstPart)
		return
	case 15:
		log.Print("They must go to jail")
		p.jailed = true
		p.square = 10
		return
	case 16:
		p.square -= 3
		log.Printf("%s They must go back three squares", firstPart)
		return
	case 17:
		p.balance -= 20
		log.Printf("%s Pays a $20 drunk in charge fine", firstPart)
		return
	case 18:
		p.balance += 100
		log.Printf("%s Collects $100 by winning a crossword competition", firstPart)
		return
	default:
		return
	}
}

func (p *Player) openCommunityChest() {
	roll := tripleDiceThrow()
	firstPart := fmt.Sprintf("P%d opens a Community Chest and rolls %d:", p.Id, roll)

	switch roll {
	case 3:
		p.balance += 200
		log.Printf("%s Collects $200 from a bank error in their favour", firstPart)
		return
	case 4:
		p.balance -= 100
		log.Printf("%s Pays $100 of hospital bills", firstPart)
		return
	case 5:
		p.balance += 10
		log.Printf("%s Collects $10 by winning second prize in a beauty contest", firstPart)
		return
	case 6:
		p.balance += 20
		log.Printf("%s Collects $20 from a tax refund", firstPart)
		return
	case 7:
		log.Printf("%s Either pay a $10 fine or take a Chance - they take a Chance!", firstPart)
		p.takeAChance()
		return
	case 8:
		p.balance -= 50
		log.Printf("%s Pays $50 of insurance premiums", firstPart)
		return
	case 9:
		p.balance += 50
		log.Printf("%s Collects $50 from a sale of stock", firstPart)
		return
	case 10:
		p.balance += 25
		log.Printf("%s Collects $25 from received interests on 7%% preference shares", firstPart)
		return
	case 11:
		p.square = 0
		p.balance += 200
		log.Printf("%s Advances to Go", firstPart)
		return
	case 12:
		p.square = 1
		log.Printf("%s Goes back to Old Kent Road", firstPart)
		return
	case 13:
		p.keys++
		log.Printf("%s Receives a Jail key", firstPart)
		return
	case 14:
		log.Printf("%s Collects $10 from each player", firstPart)
		return
	case 15:
		p.balance -= 50
		log.Printf("%s Pays $50 of doctor fees", firstPart)
		return
	case 16:
		p.balance += 100
		log.Printf("%s Collects $100 from a matured annuity", firstPart)
		return
	case 17:
		log.Print("They must go to jail")
		p.jailed = true
		p.square = 10
		return
	case 18:
		p.balance += 100
		log.Printf("%s They inherit $100", firstPart)
		return
	default:
		return
	}
}

// Returns a boolean which details whether the player has been jailed after their move
func (p *Player) move(first int, second int) bool {
	nextPos := p.square + uint8(first+second)
	nextSquare := (nextPos) % SquareCount
	p.square = nextSquare

	log.Printf("P%d rolls %d & %d - moved to square#%d", p.Id, first, second, p.square)

	if nextPos >= 40 {
		p.balance += 200
		log.Printf("P%d passed Go - they have received $200", p.Id)
	}

	switch nextSquare {
	case 2:
		fallthrough
	case 17:
		fallthrough
	case 33:
		p.openCommunityChest()
		return false
	case 7:
		fallthrough
	case 22:
		p.takeAChance()
		return false
	case 4:
		log.Printf("P%d must pay a $200 Income Tax!", p.Id)
		p.balance -= 200
		return false
	case 38:
		log.Printf("P%d must pay a $100 Super Tax!", p.Id)
		p.balance -= 100
		return false
	case 30:
		log.Printf("P%d must go to jail!", p.Id)
		p.jailed = true
		p.square = 10
		return true
	default:
		return false
	}
}

func (p *Player) PlayTurn() bool {
	doubleRollCount := 0

	for {
		first := singleDiceThrow()
		second := singleDiceThrow()

		if first == second {
			doubleRollCount++
			if doubleRollCount == 3 {
				log.Printf("P%d rolls %d & %d - their third double this turn. They go to jail!", p.Id, first, second)
				p.jailed = true
				p.square = 10
				return true
			}
		} else if first > second {
			first, second = second, first
		}

		if !p.move(first, second) || first != second {
			break
		}
	}

	return p.square != 0
}
