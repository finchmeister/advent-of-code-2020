package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

type DealtCards struct {
	player1 []int
	player2 []int
}

func FindWinningPlayersScorePt1(dealtCards DealtCards) int {
	for len(dealtCards.player1) != 0 && len(dealtCards.player2) != 0 {
		player1Card, player1RemainingCards := actionCards(dealtCards.player1)
		player2Card, player2RemainingCards := actionCards(dealtCards.player2)

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

	return computeScore(winningCards)
}

func FindWinningPlayersScorePt2(dealtCards DealtCards) int {

	player1Cards, player2Cards := playPt2(dealtCards.player1, dealtCards.player2)

	var winningCards []int
	if len(player1Cards) == 0 {
		winningCards = player2Cards
	} else {
		winningCards = player1Cards
	}

	return computeScore(winningCards)
}

func playPt2(player1Cards []int, player2Cards []int) ([]int, []int) {
	var player1PlayedCards [][]int
	var player2PlayedCards [][]int
	for len(player1Cards) != 0 && len(player2Cards) != 0 {

		player1Card, player1RemainingCards := actionCards(player1Cards)
		player2Card, player2RemainingCards := actionCards(player2Cards)

		if sameCardOrderSeen(player1Cards, player1PlayedCards) || sameCardOrderSeen(player2Cards, player2PlayedCards) {
			player1RemainingCards = append(player1RemainingCards, player1Card, player2Card)
			player1Cards = player1RemainingCards
			player2Cards = player2RemainingCards
			break
		}

		if shouldRecurse(player1Card, player1RemainingCards, player2Card, player2RemainingCards) {
			player1ReducedCards := make([]int, len(player1RemainingCards))
			player2ReducedCards := make([]int, len(player2RemainingCards))
			copy(player1ReducedCards, player1RemainingCards)
			copy(player2ReducedCards, player2RemainingCards)

			player1Cards, player2Cards = playPt2(player1ReducedCards[0:player1Card], player2ReducedCards[0:player2Card])

			if len(player1Cards) == 0 {
				player2RemainingCards = append(player2RemainingCards, player2Card, player1Card)
			} else {
				player1RemainingCards = append(player1RemainingCards, player1Card, player2Card)
			}

			player1Cards = player1RemainingCards
			player2Cards = player2RemainingCards
			continue
		}

		if player1Card > player2Card {
			player1RemainingCards = append(player1RemainingCards, player1Card, player2Card)
		} else {
			player2RemainingCards = append(player2RemainingCards, player2Card, player1Card)
		}

		player1PlayedCards = append(player1PlayedCards, player1Cards)
		player2PlayedCards = append(player2PlayedCards, player2Cards)

		player1Cards = player1RemainingCards
		player2Cards = player2RemainingCards
	}

	return player1Cards, player2Cards
}

func shouldRecurse(player1Card int, player1RemainingCards []int, player2Card int, player2RemainingCards []int) bool {
	return len(player1RemainingCards) >= player1Card && len(player2RemainingCards) >= player2Card
}

func sameCardOrderSeen(cards []int, allCardsPlayed [][]int) bool {
	for i := range allCardsPlayed {
		if reflect.DeepEqual(allCardsPlayed[i], cards) {
			return true
		}
	}

	return false
}

func actionCards(cards []int) (int, []int) {
	if len(cards) == 1 {
		return cards[0], []int{}
	}

	return cards[0], cards[1:]
}

func computeScore(cards []int) int {
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
	fmt.Println("Pt2")
	fmt.Println(FindWinningPlayersScorePt2(parse(loadFile())))
}
