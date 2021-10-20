package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	// "strconv"
)

func main(){
	// luastabung()
	// gradeNilai()
	// faktorBilangan()	
	// bilanganPrima()
	// cetakTabelPerkalian()
	// exponentiation()
	// playWithAsterisk()
	palindrome()
}

func cetakTabelPerkalian(){
	var inputUser int
	print("Masukkan nilai: ")
	fmt.Scan(&inputUser)
	if(inputUser<1||inputUser>30){
		println("invalid input")
	}else{
		for i:=1;i<=inputUser;i++{
			for j:=1;j<=9;j++{
				print(i*j)
				print(" ")
			}
			println()
		}
	}
}

func playWithAsterisk(){
	
	var inputUser int
	print("Masukkan nilai: ")
	fmt.Scan(&inputUser)

	for i:=1;i<=inputUser;i++{
		for j:=inputUser-i;j>=1;j--{
			print(" ")
		}
		for k:=1;k<=i;k++{
			print("* ")
		}
		println()
	}
}


func exponentiation(){
	
	var x,n float64
	print("base: ")
	fmt.Scan(&x)
	print("exponent: ")
	fmt.Scan(&n)

	sum:=math.Pow(x,n)
	// s:=strconv.Itoa(int(sum))
	print(int(sum))
}

func palindrome(){
	
	print("Masukkan kata: ")
	baca := bufio.NewScanner(os.Stdin)
	baca.Scan() 
	line := baca.Text()
	// fmt.Printf("Input was: %q\n", line)
	for i:=0;  i<len(line);i++{
		fmt.Printf("%c\n",line[i])
	}

	kebalik:=len(line)
	for j := kebalik; j >0; j-- {
		fmt.Printf("%c\n",line[j-1])
	}
	
}

func bilanganPrima(){
	var inputUser int
	print("Masukkan nilai: ")
	fmt.Scan(&inputUser)

	
	var m,flag int
	m=0
	flag=0
	m=inputUser/2
	
	if(inputUser<=1){
		println("false")
		flag=1
	}else{
	for i:=2;i<=m;i++{
		if inputUser%i==0{
			println("false")
			flag=1
			break
		}
	}}

	if(flag==0){
			println("Bilangan prima")
	}

}

func faktorBilangan(){
	var inputUser int
	print("Masukkan nilai: ")
	fmt.Scan(&inputUser)
	for i:=1;i<=inputUser;i++{
		if(inputUser%i==0){
			println(i)
		}
	}
}

func gradeNilai(){
	var studentScore int
	var namaSiswa string
	var keteranganNilai string

	print("Masukkan nama siswa: ")
	fmt.Scan(&namaSiswa)

	print("Masukkan nilai "+namaSiswa+": ")
	fmt.Scan(&studentScore)

	if studentScore >=0 && studentScore<=34{
		keteranganNilai="E"
	}else if studentScore >=35 && studentScore<=49{
		keteranganNilai="D"
	}else if studentScore >=50 && studentScore<=64{
		keteranganNilai="C"
	}else if studentScore >=65 && studentScore<=79{
		keteranganNilai="B"
	}else if studentScore >=80 && studentScore<=100{
		keteranganNilai="A"
	}else{
		keteranganNilai="tidak valid"
	}

	// nilai:=strconv.Itoa(studentScore)
	
	fmt.Print("Nilai dari ",namaSiswa," adalah ",studentScore," bernilai ",keteranganNilai)
	
}

func luastabung(){
	var T,r,pi float32
	pi=3.14

	print("Masukkan tinggi: ")
	fmt.Scan(&T)
	print("Masukkan radius: ")
	fmt.Scan(&r)
	
	var selimut float32
	selimut=2*pi*r*(r+T)

	fmt.Print(selimut)
}