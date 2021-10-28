package main

import "fmt"

func remove(slice []int, s int) {
	slice = append(slice[:s], slice[s+1:]...)
}

func DragonOfLoowater(dragonHead, knightHeight []int) {

	// your code here
	if len(dragonHead) > len(knightHeight) {
		fmt.Println("knight fall")
		return
	}

	var minTotHeight, tempH int
	var isKilled bool
	var index int

	for _, diameter := range dragonHead {
		tempH = 0
		isKilled = false
		for i, height := range knightHeight {
			if height >= diameter && height < tempH {
				minTotHeight += height - tempH
				tempH = height
				index = i
				continue
			} else if isKilled {
				continue
			} else if height >= diameter {
				minTotHeight += height
				tempH = height
				index = i
				isKilled = true
			}
		}

		if !isKilled {
			fmt.Println("knight fall")
			return
		}

		remove(knightHeight, index)
	}

	fmt.Println(minTotHeight)
	return
}

func main() {

	DragonOfLoowater([]int{5, 4}, []int{7, 8, 4}) // 11

	DragonOfLoowater([]int{5, 10}, []int{5}) // knight fall

	DragonOfLoowater([]int{7, 2}, []int{4, 3, 1, 2}) // knight fall

	DragonOfLoowater([]int{7, 2}, []int{2, 1, 8, 5}) // 10

	DragonOfLoowater([]int{10, 9}, []int{10, 1, 8, 5}) // knight fall
}