package class2

import "fmt"

var arrayExample = [7]int{1,2,3,4}
var sliceExample = []int{}

func ShowExamples() {
	//dodawanie elementow do slice'ów
	sliceExample = append(sliceExample,9,3,4,5,7,4,2)
	fmt.Println(sliceExample)
	sliceExample = append(sliceExample[:2],sliceExample[3:]...)
	fmt.Println(sliceExample)

	str:= "string"
	str = string(append([]byte(str),byte(63)))
	fmt.Println(str)
}