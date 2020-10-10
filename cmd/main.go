package main

import (
	"log"

	"goCourse/class3"
)

func main() {
	h := class3.NewHasher("in", "out")
	err := h.ReadLines()
	handleError(err)
	err = h.HashLines()
	handleError(err)
	err = h.SaveToFile()
	handleError(err)
}
func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
