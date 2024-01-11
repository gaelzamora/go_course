package main

import (
	"fmt"
	"github.com/gaelzamora/go_course/variables"
)

func main() {
	estado, texto := variables.ToText(5)
	fmt.Println(estado)
	fmt.Println(texto)
}