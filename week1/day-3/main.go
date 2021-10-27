package main

import (
	"fmt"
	"math"
	"strconv"
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
	// fmt.Println(PairSum([]int{8, 5, 1, 2, 3, 4}, 6)) // [1, 3]
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
	input:="12341235"
	a:=[]int{}
	for i := 0; i < len(input); i++ {
		temp,_:=strconv.Atoi( input[i:i+1])
		a = append(a, temp)
	}
	b:=a
	for i := 0; i < len(a); i++ {
		for j := i+1; j < len(a); j++ {
			if a[i]==a[j]{
				b[i]=0
				b[j]=0
			}
		}
	}
	for i := 0; i < len(b); i++ {
		if b[i]==0{
			copy(b[i:], b[i+1:]) // Shift a[i+1:] left one index.
			b[len(b)-1] = 0     // Erase last element (write zero value).
			b = b[:len(b)-1]     // Truncate slice.
			i=-1
		}
	}
	fmt.Printf("%v",b)
}
func pairWithTargetSum(){
	a:=[]int{7,1,2,3,4,5,6}
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
func PairSum(arr []int, target int) []int {
	l := 0
	r := len(arr) - 1
	sum := arr[l] + arr[r]
	for sum != target {
		if sum > target {
			r -= 1
		} else {
			l += 1
		}
		sum = arr[l] + arr[r]
	}
	return []int{l, r}
}