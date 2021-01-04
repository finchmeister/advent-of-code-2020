package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type DealtCards struct {
	player1 []int
	player2 []int
}

func FindWinningPlayersScorePt1(dealtCards DealtCards) int {
	for len(dealtCards.player1) != 0 && len(dealtCards.player2) != 0 {
		player1Card, player1RemainingCards := actionCardsPt1(dealtCards.player1)
		player2Card, player2RemainingCards := actionCardsPt1(dealtCards.player2)

		if player1Card > player2Card {
			player1RemainingCards = append(player1RemainingCards, player1Card, player2Card)
		} else {
			player2RemainingCards = append(player2RemainingCards, player2Card, player1Card)
		}

		dealtCards.player1 = player1RemainingCards
		dealtCards.player2 = player2RemainingCards
	}

	var winningCards []int
	if len(dealtCards.player1) == 0 {
		winningCards = dealtCards.player2
	} else {
		winningCards = dealtCards.player1
	}

	return computeScorePt1(winningCards)
}

func actionCardsPt1(cards []int) (int, []int) {
	if len(cards) == 1 {
		return cards[0], []int{}
	}

	return cards[0], cards[1:]
}

func computeScorePt1(cards []int) int {
	score := 0
	for i, card := range cards {
		score = score + (len(cards)-i)*card
	}

	return score
}

func parse(input string) DealtCards {
	rawCardSections := strings.Split(input, "\n\n")

	return DealtCards{
		parseCardSection(rawCardSections[0]),
		parseCardSection(rawCardSections[1]),
	}
}

func parseCardSection(rawCardSection string) []int {

	var cards []int
	for _, row := range strings.Split(rawCardSection, "\n") {
		card, err := strconv.Atoi(row)
		if err != nil {
			continue
		}
		cards = append(cards, card)
	}

	return cards
}

func loadFile() string {
	data, err := ioutil.ReadFile("22_input.txt")
	if err != nil {
		panic(err)
	}

	return string(data)
}

func main() {
	fmt.Println("Pt1")
	fmt.Println(FindWinningPlayersScorePt1(parse(loadFile())))
}
