package monopoly

type PlayerTurn struct {
	player *Player
	next   *PlayerTurn
}

type TurnQueue struct {
	head *PlayerTurn
	tail *PlayerTurn
}

func (q *TurnQueue) Push(p *Player) {
	q.tail.next = &PlayerTurn{
		player: p,
		next:   nil,
	}
	q.tail = q.tail.next
}

func (q *TurnQueue) Pop() *Player {
	player := q.head.player
	q.head = q.head.next
	return player
}
