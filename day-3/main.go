package main

import (
	"fmt"
	"math"
)
func main(){
	// nomor 1
	// bilanganPrima()

	// nomor 2
	// exponentiation()

	// nomor 3
	// arrayMerge()

	// nomor 4
	// angkaMunculSekali()
	
	// nomor 5
	pairWithTargetSum()
}
func bilanganPrima(){
	var inputUser int
	print("Masukkan nilai: ")
	fmt.Scan(&inputUser)
	var m,flag int
	m=0
	flag=0
	m=int(math.Sqrt(float64(inputUser)))
	if(inputUser<=1){
		println("not prime")
		flag=1
	}else{
	for i:=2;i<=m;i++{
		if inputUser%i==0{
			println("not prime")
			flag=1
			break
		}
	}}
	if(flag==0){
			println("Bilangan prima")
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

func arrayMerge(){
	a:=[]string{"king","devil jin","akuma"}
	b:=[]string{"eddie","steve","geese"}
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b); j++ {
			if a[i]==b[j] {
				copy(b[i:], b[i+1:]) // Shift a[i+1:] left one index.
				b[len(b)-1] = ""     // Erase last element (write zero value).
				b = b[:len(b)-1]     // Truncate slice.
			}
		}
	}
	c:=append(a,b...)
	fmt.Println(c)
}

func angkaMunculSekali(){
	a:=[]int{1,2,3,4,5,1,2,3}

	b:=[]int{}
	for i := 0; i < len(a); i++ {
		for j := i+1; j < len(a); j++ {
			if a[i]==a[j]{
				b=append(b, a[i])
			}
		}
	}

	fmt.Printf("%v",b)
}

func pairWithTargetSum(){
	a:=[]int{1,2,3,4,5,6}
	target:=7

	for i := 0; i < len(a); i++ {
		for j := i+1; j < len(a); j++ {
			if a[i]+a[j]==target {
				c:=[]int{i,j}
				fmt.Printf("%v",c)
				break
			}
		}
	}
}