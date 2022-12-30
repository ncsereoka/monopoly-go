package monopoly

import (
	"fmt"
	"log"
	"math/rand"
)

var PlayerMap = map[int]*Player{}

type Player struct {
	Id         int
	balance    int
	square     uint8
	jailed     bool
	keys       uint8
	stations   uint8
	utilities  uint8
	browns     uint8
	lightBlues uint8
	pinks      uint8
	oranges    uint8
	reds       uint8
	yellows    uint8
	greens     uint8
	darkBlues  uint8
	houses     uint8
	hotels     uint8
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

func (p *Player) pay(fee int) {
	p.balance -= fee
}

func (p *Player) beAssessedForStreetRepairs() {
	sum := p.houses*40 + p.hotels*115
	log.Printf("P%d owns %d house(s) and %d hotel(s) - they must pay $%d", p.Id, p.houses, p.hotels, sum)
}

func (p *Player) makeGeneralRepairs() {
	sum := p.houses*25 + p.hotels*100
	log.Printf("P%d owns %d house(s) and %d hotel(s) - they must pay $%d", p.Id, p.houses, p.hotels, sum)
}

func (p *Player) takeAChance() {
	roll := tripleDiceThrow()
	firstPart := fmt.Sprintf("P%d takes a Chance and rolls %d:", p.Id, roll)

	switch roll {
	case 3:
		p.pay(15)
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
		p.pay(150)
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
		p.beAssessedForStreetRepairs()
		return
	case 11:
		log.Printf("%s Makes general repairs on all of their buildings", firstPart)
		p.makeGeneralRepairs()
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
		p.pay(20)
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
		p.pay(100)
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
		p.pay(50)
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
		p.pay(50)
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

	// log.Printf("P%d rolls %d & %d", p.Id, first, second)
	firstPart := fmt.Sprintf("P%d moved to square#%d:", p.Id, p.square)

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
		log.Print(firstPart, " Community Chest!")
		p.openCommunityChest()
		return false
	case 7:
		fallthrough
	case 22:
		fallthrough
	case 36:
		log.Print(firstPart, " Chance!")
		p.takeAChance()
		return false
	case 4:
		log.Printf("%s P%d must pay a $200 Income Tax!", firstPart, p.Id)
		p.pay(200)
		return false
	case 38:
		log.Printf("%s P%d must pay a $100 Super Tax!", firstPart, p.Id)
		p.pay(100)
		return false
	case 30:
		log.Printf("%s P%d must go to Jail!", firstPart, p.Id)
		p.jailed = true
		p.square = 10
		return true
	case 12:
		fallthrough
	case 28:
		utility := PropertyMap[int(p.square)]
		log.Printf("%s %s", firstPart, utility.name)
		if utility.owner == 0 {
			if utility.price < p.balance {
				p.pay(utility.price)
				p.utilities++
				utility.owner = int(p.Id)
				log.Printf("P%d bought %s", p.Id, utility.name)
			}
		} else {
			roll := first + second
			owner := PlayerMap[utility.owner]

			var fee int
			if owner.utilities == 1 {
				fee = roll * 4
			} else {
				fee = roll * 10
			}
			log.Printf("P%d owes P%d $%d", p.Id, owner.Id, fee)
			p.pay(fee)
			owner.balance += fee
		}
		return false
	case 5:
		fallthrough
	case 15:
		fallthrough
	case 25:
		fallthrough
	case 35:
		station := PropertyMap[int(p.square)]
		log.Printf("%s %s Station", firstPart, station.name)
		if station.owner == 0 {
			if station.price < p.balance {
				p.pay(station.price)
				p.stations++
				station.owner = int(p.Id)
				log.Printf("P%d bought %s", p.Id, station.name)
			}
		} else {
			owner := PlayerMap[station.owner]
			var fee int
			switch owner.stations {
			case 1:
				fee = 25
			case 2:
				fee = 50
			case 3:
				fee = 100
			case 4:
				fee = 200
			}

			log.Printf("P%d owes P%d $%d", p.Id, owner.Id, fee)
			p.pay(fee)
			owner.balance += fee
		}
		return false
	case 10:
		log.Printf("%s Just visiting", firstPart)
		return false
	case 20:
		log.Printf("%s Free Parking", firstPart)
		return false
	case 0:
		log.Printf("%s Go", firstPart)
		return false
	default:
		log.Printf("%s %s", firstPart, PropertyMap[int(p.square)].name)
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

	return p.balance < 10000
}
