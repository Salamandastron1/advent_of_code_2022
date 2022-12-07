package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const win, draw, loss int = 6, 3, 0

type handsign struct {
	name     string
	elfID    string
	playerID string
	reward   int
}

var rock handsign = handsign{"rock", "A", "X", 1}
var paper handsign = handsign{"paper", "B", "Y", 2}
var scissors handsign = handsign{"scissors", "C", "Z", 3}
var hands []handsign = []handsign{rock, paper, scissors}

func main() {
	totalScore := 0
	f, err := os.Open("strategy.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		handSigns := strings.Split(line, " ")
		fmt.Println(handSigns)
		totalScore += calculateScore(handSigns[0], handSigns[1])
	}
	fmt.Println(totalScore)
}

func calculateScore(elfHandsign, myHandsign string) int {
	elf := handsign{}
	me := handsign{}
	for _, v := range hands {
		if myHandsign == v.playerID {
			me = v
		}
		if elfHandsign == v.elfID {
			elf = v
		}
	}
	fmt.Println("Elf hand", elf)
	fmt.Println("My hand", me)
	fmt.Println(battleOutcome(elf, me))
	return battleOutcome(elf, me)
}

func battleOutcome(elfHand, myHand handsign) int {
	switch {
	case elfHand.name == "scissors" && myHand.name == "paper":
		return loss + myHand.reward
	case elfHand.name == "rock" && myHand.name == "paper":
		return win + myHand.reward
	case elfHand.name == "paper" && myHand.name == "rock":
		return loss + myHand.reward
	case elfHand.name == "scissors" && myHand.name == "rock":
		return win + myHand.reward
	case elfHand.name == "rock" && myHand.name == "scissors":
		return loss + myHand.reward
	case elfHand.name == "paper" && myHand.name == "scissors":
		return win + myHand.reward
	default:
		return draw + myHand.reward
	}
}
