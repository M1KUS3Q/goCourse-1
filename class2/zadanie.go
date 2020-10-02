package class2

import (
	"flag"
	"fmt"
	"strconv"
)

func Apples() {
	Japka := flag.Int("japka", 0, "Sumuje podane argumenty")
	flag.Parse()
	fmt.Println("Japłek jest " + strconv.Itoa(*Japka) + "!")
}
