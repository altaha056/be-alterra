package main

import (
	"fmt"
)

func sumEl(arr []int) int {

	// your code here
	var sum int
	for _, v := range arr {
		sum += v
	}

	return sum
}

func canOut(hand, deck []int) bool {
	if len(hand) != len(deck) {
		return false
	}

	for _, v := range hand {
		for _, w := range deck {
			if v == w {
				return true
			}
		}
	}

	return false
}

func playingDomino(cards [][]int, deck []int) interface{} {

	// your code here
	var iPossOutCard []int

	for i, v := range cards {
		if canOut(v, deck) {
			iPossOutCard = append(iPossOutCard, i)
		}
	}

	if len(iPossOutCard) == 0 {
		return "tutup kartu"
	}

	outCard := cards[iPossOutCard[0]]

	for _, v := range iPossOutCard {
		if sumEl(cards[v]) > sumEl(outCard) {
			outCard = cards[v]
		}
	}

	return outCard
}

func main() {

	fmt.Println(playingDomino([][]int{{6, 5}, {3, 4}, {2, 1}, {3, 3}}, []int{4, 3}))

	// [3, 4]

	fmt.Println(playingDomino([][]int{{6, 5}, {3, 3}, {3, 4}, {2, 1}}, []int{3, 6}))

	// [6 5]

	fmt.Println(playingDomino([][]int{{6, 6}, {2, 4}, {3, 6}}, []int{5, 1}))

	// “tutup kartu”

}