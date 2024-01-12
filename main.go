package main

import (
	"fmt"
	"github.com/gaelzamora/go_course/variables"
)

func main() {
	state, number := variables.GetMult(10)
	fmt.Println(state)
	fmt.Println(number)
	fmt.Println(variables.GetSay("Gael"))
}