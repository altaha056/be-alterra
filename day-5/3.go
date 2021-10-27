package main
import (
	"fmt"
	"math"
)
func primaSegiEmpat(high, wide, start int) {
	fmt.Printf("\n\n")
	var sum int
	mapOfPrime := mapOfPrime(high, wide, start)
	j := 1
	for i := 1; i <= wide; i++ {
		for j <= i*high {
			fmt.Printf("%d\t", mapOfPrime[j])
			j++
		}
		fmt.Printf("\n\n\n")
	}
	for _, v := range mapOfPrime {
		sum += v
	}
	println(sum)
	fmt.Printf("\n")
}

func main() {
	primaSegiEmpat(2, 3, 13)
	/*

	   17 19

	   23 29

	   31 37

	   156

	*/
	primaSegiEmpat(5, 2, 1)
	/*
	   2  3  5  7 11

	   13 17 19 23 29

	   129

	*/
}

func mapOfPrime(high, wide, start int) map[int]int {

	var index int
	indexToPrime := make(map[int]int)

	i := start + 1

	for i >= start+1 {
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
			indexToPrime[index] = i
		}

		if index == high*wide {
			break
		}

		i++
	}

	return indexToPrime
}