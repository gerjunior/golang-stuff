package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

func main() {
	fmt.Print("Hello, ")
	fmt.Println("World!")
	fmt.Printf("Another program in %s\n", "Go")

	const PI float64 = 3.1415
	var raio = 3.2
	area := PI * math.Pow(raio, 2)
	fmt.Println(area)

	// const (
	// 	a = 1
	// 	b = 2
	// )

	// var (
	// 	c = 3
	// 	d = 4
	// )

	// fmt.Println(a, b, c, d)

	var e, f bool = true, false
	g, h, i := 2, false, "epa!"
	fmt.Println(e, f, g, h, i)

	// x := 3.1415
	// xs := fmt.Sprint(x)
	// fmt.Println("The value of x is " + xs)

	// fmt.Printf("The value of x is %.2f.", x)
	a := 2
	b := 1.9999
	c := false
	d := "hello!"
	fmt.Printf("\n%d %f %.1f %t %s\n", a, b, b, c, d)

	fmt.Println("Literal int is ", reflect.TypeOf(a))

	var unicodeChar rune = 'a'
	fmt.Println("unicode char a: ", unicodeChar)

	s2 := `Hello
	World
	!`

	fmt.Println(s2)

	// default types
	var intenger int
	var floater float64
	var boolean bool
	var stringV string
	var pointer *int

	// 0 0 false "" <nil>
	fmt.Printf("%v %v %v %q %v", intenger, floater, boolean, stringV, pointer)

	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9

	notaFinal := int(nota)
	// 6
	fmt.Println(notaFinal)

	// int to string
	fmt.Println("Teste " + fmt.Sprint(97))
	fmt.Println("Teste " + strconv.Itoa(97))

	// string to int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	booleanValue, _ := strconv.ParseBool("true") // only 1 and "true" are true
	if booleanValue {
		fmt.Println("True")
	}
}
