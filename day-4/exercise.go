package main

import "fmt"

// import "fmt"
func main(){
	defer fmt.Print("done")
	// println(penjumlahan(5,5,5,2))
	// println(penjumlahan(5,88,5,2))
	// kataSambutan:=func(nama string){
	// 	println("Halo selamat datang",nama)
	// }

	// summ:=func(data...int){
	// 	a:=0
	// 	for _,v:=range data{
	// 		a=a+v
	// 	}
	// 	println(a)
	// }
	// summ(2,3,4,5,6)

	// kataSambutan("altaha")
	// count:=counter()
	// fmt.Println(count())
	a:="altaha"
	b:=&a
	*b="alta"
	fmt.Println(b,a)
}
func penjumlahan(data...int)int{
	sum:=0
	for _,v:=range data{
		sum+=v
	}
	return sum
}
func counter() func()int{
	count:=0
	return(func() int {
		count+=1
		return(count)
	})
}

