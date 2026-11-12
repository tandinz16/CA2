package rps

import (
	"math/rand"
)

const (
	ROCK         = 0 // beats scissors
	PAPER        = 1 // beats rock
	SCISSORS     = 2 // beats paper
	PLAYERWINS   = 1 // player wins
	COMPUTERWINS = 2 // computer wins
	DRAW         = 3 // draw
)

type Round struct {
	Winner         int
	ComputerChoice string
	RoundResult    string
}

// returns game result, computer choice, round result
func PlayRound(playerValue int) Round {

	// get random value for computer choice
	computerValue := rand.Intn(3)

	computerChoice := ""
	roundResult := ""
	winner := 0

	var result Round

	// check what computer chooses
	switch computerValue {
	case ROCK:
		computerChoice = "Computer chose ROCK"

	case PAPER:
		computerChoice = "Computer chose PAPER"

	case SCISSORS:
		computerChoice = "Computer chose SCISSORS"

	default:
		return result
	}

	// check winner
	if playerValue == computerValue {
		roundResult = "It's a draw"
		winner = DRAW

	} else if playerValue == (computerValue+1)%3 {
		roundResult = "Player wins!"
		winner = PLAYERWINS

	} else {
		roundResult = "Computer wins!"
		winner = COMPUTERWINS
	}

	result = Round{
		Winner:         winner,
		ComputerChoice: computerChoice,
		RoundResult:    roundResult,
	}

	return result
}
