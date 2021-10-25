package main

import (
	"fmt"
	"math"
)

func Compare(a, b string) string {

	maxLen := int(math.Max(float64(len(a)), float64(len(b))))
	shortLen := int(math.Min(float64(len(a)), float64(len(b))))
	longStr := b
	shortStr := a

	if maxLen == len(a) {
		longStr = a
		shortStr = b
	}

	var initStr, same string
	var temp int

	for i := range shortStr {
		for j := range longStr {
			temp = 0
			same = ""

			for int(i+temp) < shortLen && int(j+temp) < maxLen && shortStr[i+temp] == longStr[j+temp] {
				same += string(longStr[j+temp])
				temp += 1
			}

			if len(same) > len(initStr) {
				initStr = same
			}
		}
	}
	return initStr
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI")) // AKA
	fmt.Println(Compare("KANGOORO", "KANG")) // KANG
	fmt.Println(Compare("KI", "KIJANG")) // KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU
	fmt.Println(Compare("ILALANG", "ILA")) // ILA
	fmt.Println(Compare("SEPATU", "EPA")) // EPA
}