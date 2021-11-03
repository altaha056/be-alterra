package main
import (
	"fmt"
	"strings"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func countPrintChar(sentence string, index int, isCapital bool) {
	N := 97
	if isCapital {
		N = 65
	}

	theChar := string(rune(index + N))
	countFreq := strings.Count(sentence, theChar)
	if countFreq != 0 {
		fmt.Printf("%s : %d \n", theChar, countFreq)
	}
	wg.Done()
}

func main() {
	start := time.Now()
	sentence := "LLLLLL ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"

	for i := 0; i < 26; i++ {
		wg.Add(2)
		go countPrintChar(sentence, i, false) // Counting the non-Capital character
		go countPrintChar(sentence, i, true)  // Counting the capital character
	}
	wg.Wait()
	fmt.Println(time.Since(start))
}