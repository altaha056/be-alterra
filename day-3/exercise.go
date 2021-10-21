package main

import (
	"fmt"
)

func main(){
	jumlah:=0
	print("masukkan jumlah: ")
	fmt.Scan(jumlah)
	
	var anggota[]string
	
	for i := 0; i < jumlah; i++ {
		print("anggota ke-",i,": ")
		fmt.Scan(anggota[i])

	}
	
	fmt.Println(anggota)
}