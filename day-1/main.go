package main

import (
	testPackage2 "day1/testpackage2"
	"fmt"
)

func main(){
	hello()
	fmt.Print("Hello everybody")
	testPackage2.Greet()
}

func hello(){
	var i float32
	i=32.4
	fmt.Printf("%v, %T\n",i,i)
	var j int
	j=int(i)
	fmt.Printf("%v, %T\n",j,j)
	
}