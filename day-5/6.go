package main

import (
	"fmt"
	"strconv"
)

func FindMax(arr []int) int {

	// your code here
	if len(arr) == 0 {
		return 0
	}

	max := arr[0]

	for _, v := range arr {
		if v > max {
			max = v
		}
	}

	return max

}

func MaximumBuyProduct(money int, productPrice []int) {

	// your code here
	var maxProductBought []int

	for i := 0; i < len(productPrice); i++ {
		bought := productPrice[i]
		productBought := 1

		if bought > money {
			bought -= productPrice[i]
			productBought--
		} else {
			for j := i + 1; j < len(productPrice); j++ {
				bought += productPrice[j]
				productBought++
				if bought > money {
					bought -= productPrice[j]
					productBought--
				}
			}
			maxProductBought = append(maxProductBought, productBought)
		}
	}

	fmt.Println(strconv.Itoa(FindMax(maxProductBought)))
}

func main() {

	MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000}) // 3

	MaximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000}) // 4

	MaximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000}) // 4

	MaximumBuyProduct(4000, []int{7500, 3000, 2500, 2000}) // 1

	MaximumBuyProduct(0, []int{10000, 30000}) // 0

}