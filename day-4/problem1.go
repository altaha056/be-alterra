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

	var comStr, theSame string
	var temp int

	for i := range shortStr {
		for j := range longStr {
			temp = 0
			theSame = ""

			for int(i+temp) < shortLen && int(j+temp) < maxLen && shortStr[i+temp] == longStr[j+temp] {
				theSame += string(longStr[j+temp])
				temp += 1
			}

			if len(theSame) > len(comStr) {
				comStr = theSame
			}
		}
	}
	return comStr
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI")) // AKA
	fmt.Println(Compare("KANGOORO", "KANG")) // KANG
	fmt.Println(Compare("KI", "KIJANG")) // KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU
	fmt.Println(Compare("ILALANG", "ILA")) // ILA
	fmt.Println(Compare("SEPATU", "EPA")) // EPA
	fmt.Println(Compare("ABCD", "KFAN")) //
}