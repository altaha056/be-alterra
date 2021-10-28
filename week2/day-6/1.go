package main
import "fmt"
func main(){
	calculate(6,6,14)
}

func calculate(a,b,c int)(x,y,z int){
	q:=(a+b+c)/3
	for i := 1; i <= q; i++ {
		for j := 1; j <=q; j++ {
			for k := 1; k <= q; k++ {
				if i+j+k==a {
					if i*j*k==b{
						if i*i+j*j+k*k==c{
							fmt.Println(i,j,k)
							return
						}
					}
				}
			}
		}
	}
	fmt.Println("no solution found")
	return
}