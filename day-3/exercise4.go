package main
import "fmt"
func main(){
	// h:=0
	// print("masukkan jam: ")
	// fmt.Scan(&h)
	// jam(h)

	var a,t,s int
	print("masukkan alas: ")
	fmt.Scanln(&a)
	
	print("masukkan tinggi: ")
	fmt.Scanln(&t)

	print("masukkan panjang sisi: ")
	fmt.Scanln(&s)

	segitiga(a,t,s)
}

func segitiga(a,t,s int)(int,int){
	luas:=a*t/2
	keliling:=s*3
	println("Luas segitiga adalah:",luas)
	println("Keliling segitiga adalah:",keliling)
	return luas,keliling


}

func jam(hour int){
	if(hour<=12&&hour>0){
		fmt.Println("am")
	}else if(hour>12&&hour<=24){
		fmt.Println("pm")
	}else{
		fmt.Println("not valid")
	}

}