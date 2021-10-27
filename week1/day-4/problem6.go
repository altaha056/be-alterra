package main
import (
	"fmt"
	"strings"
)

type student struct {
	name string

	nameEncode string

	score int
}

type Chiper interface {
	Encode() string

	Decode() string
}

func (s *student) Encode() string {

	var nameEncode = ""

	Encoding := func(r rune) rune {

		s := int(r) + 3
		if s > 'z' {
			return rune(s - 26)
		} else if s < 'a' {
			return rune(s + 26)
		}
		return rune(s)
	}

	nameEncode = strings.Map(Encoding, s.name)

	// your code here

	return nameEncode

}

func (s *student) Decode() string {

	var nameDecode = ""

	Decoding := func(r rune) rune {

		s := int(r) - 3
		if s > 'z' {
			return rune(s - 26)
		} else if s < 'a' {
			return rune(s + 26)
		}
		return rune(s)
	}

	nameDecode = strings.Map(Decoding, s.nameEncode)

	// your code here

	return nameDecode

}

func main() {

	var menu int

	var a = student{}

	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu ? ")

	fmt.Scan(&menu)

	if menu == 1 {

		fmt.Print("\nInput Student’s Name : ")

		fmt.Scan(&a.name)

		fmt.Print("\nEncode of Student’s Name " + a.name + " is : " + c.Encode())

	} else if menu == 2 {

		fmt.Print("\nInput Student’s Encode Name : ")

		fmt.Scan(&a.nameEncode)

		fmt.Print("\nDecode of Student’s Name " + a.nameEncode + " is : " + c.Decode())

	} else {

		fmt.Println("Wrong input name menu")

	}

}