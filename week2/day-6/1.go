package main
import "fmt"
func main(){
	calculate(3,1,3)
}

func calculate(a,b,c int)(x,y,z int){
	q:=(a+b+c)/3
	for i := 1; i <= q; i++ {
		for j := 1; j <=q; j++ {
			for k := 1; k <= q; k++ {
				if i+j+k==a&&i*j*k==b&&i*i+j*j+k*k==c {
					fmt.Println(i,j,k)
					return
				}
			}
		}
	}
	fmt.Println("no solution found")
	return
}