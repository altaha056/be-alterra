package main
import "fmt"

type Student struct {
	name  []string
	score []int
}

func (s Student) Avarage() float64 {
	myMap := make(map[string]int)
	for _, v := range s.score {
		myMap[""] += v
	}

	return float64(myMap[""] / len(s.score))
}

func (s Student) Min() (min int, name string) {
	myMap := map[string]int{
		"": s.score[0],
	}

	var index int

	for i, v := range s.score {
		if myMap[""] > v {
			myMap[""] = v
			index = i
		}
	}

	return myMap[""], s.name[index]
}

func (s Student) Max() (max int, name string) {
	myMap := map[string]int{
		"": s.score[0],
	}

	var index int

	for i, v := range s.score {
		if myMap[""] < v {
			myMap[""] = v
			index = i
		}
	}

	return myMap[""], s.name[index]
}

func main() {

	var a = Student{}
	for i := 0; i < 6; i++ {
		var name string
		fmt.Print("Input " + string(i) + " Student’s Name : ")
		fmt.Scan(&name)
		a.name = append(a.name, name)
		var score int
		fmt.Print("Input " + name + " Score : ")
		fmt.Scan(&score)
		a.score = append(a.score, score)

	}
	fmt.Println("\n\nAvarage Score Students is", a.Avarage())
	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Students is : "+nameMax+" (", scoreMax, ")")
	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Students is : "+nameMin+" (", scoreMin, ")")
}