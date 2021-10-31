package main

import (
    "fmt"
)

func firstOccurrence(s string, x string) int32 {
	// Write your code here
	pos := 0
	check := 0
	for i := 0; i < len(s); i++ {

		if s[i] == x[check] || x[check] == '*' {
			check++
			if check == len(x) {
				return int32(pos)
			}

		} else if s[i] != x[check] {
			check = 0
			pos = i
		}

	}

	return -1
}

func main() {
    s := "xabcdey"
    x := "ab*de"

    fmt.Println(firstOccurrence(s,x))
}
