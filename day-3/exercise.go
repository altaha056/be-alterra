package main

import (
	"fmt"
)

func main(){
	jumlah:=0
	print("masukkan jumlah: ")
	fmt.Scan(&jumlah)

 	a := make([]string, jumlah)
	
	for i := 0; i < jumlah; i++ {
		print("anggota ke-",i+1,": ")
		fmt.Scan(&a[i])

	}
	
	for j:=0;j<jumlah;j++{
		fmt.Println(a[j])
	}
}