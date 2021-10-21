package main
import "fmt"
func main(){
	murid:=0
	print("Jumlah murid: ")
	fmt.Scan(&murid)
	// banyakAnggota:=murid
	var nama=[]string{}

	for i:=1;i<=murid;i++{
		print("Nama murid ke-",i)
		fmt.Scan(&nama[i])
	}
	
	// var colors=[]string{"red","green","blue","orange","tosca","brown"}
	// colors=append(colors, "purple")
	// copiedColors:=make([]string,4)
	// copy(copiedColors,colors)
	// fmt.Println(copiedColors)

	// for i:=0;i<len(colors);i++{
	// 	fmt.Println(colors[i])
	// }

}