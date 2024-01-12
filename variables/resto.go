package variables

import (
	"fmt"
	"time"
	"strconv"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	Nombre = "Pedro"
	Estado = true
	Sueldo = 1577.66
	Fecha = time.Now()
	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}

func ToText(number int) (bool, string) {
	text := strconv.Itoa(number)
	return true, text
}

func GetMult(number int) (bool, int) {
	mult := number * number
	return true, mult
}

func GetSay(name string) (string) {
	text := "Hola espero estes bien:) "
	return text+name
}	