package main
import (
	"fmt"
	"math"
)
func primeX(number int) int {
	var index int
	i := 2
	for i >= 2 {
		isPrime := true
		sqrtn := int(math.Sqrt(float64(i)))
		j := 2
		for j <= sqrtn {
			if i%j == 0 {
				isPrime = false
				j = i
			}
			j++
		}
		if isPrime {
			index++
		}

		if index == number {
			return i 
		}

		i++
	}
	return i 
}
func main() {
	fmt.Println(primeX(1)) // 2
	fmt.Println(primeX(5)) // 11
	fmt.Println(primeX(8)) // 19
	fmt.Println(primeX(9)) // 23
	fmt.Println(primeX(10)) // 29
}