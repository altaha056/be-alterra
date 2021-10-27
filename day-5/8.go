package main

import (
	"fmt"
	"strconv"
)

type pair interface{}

type appear struct {
	item  string
	count int
}

func MostAppearItem(items []string) []pair {

	// your code here

	itemToCount := make(map[string]int)

	for _, v := range items {
		itemToCount[v]++
	}

	sortedResult := make([]appear, len(itemToCount))

	index := 0
	for k, v := range itemToCount {
		sortedResult[index].item = k
		sortedResult[index].count = v
		index++
	}

	for i := range sortedResult {
		for j := i + 1; j < len(sortedResult); j++ {
			if sortedResult[i].count > sortedResult[j].count {
				temp := sortedResult[i]
				sortedResult[i] = sortedResult[j]
				sortedResult[j] = temp
			}
		}
	}

	var output []pair

	for _, v := range sortedResult {
		output = append(output, v.item+"->"+strconv.Itoa(v.count))
	}
	// return sortedResult

	return output
}

func main() {

	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))

	// golang->1 ruby->2 js->4

	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))

	// C->1 D->2 B->3 A->4

	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))

	// football->1 basketball->1 tenis->1

}