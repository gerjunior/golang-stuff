package main

type car struct {
	name  string
	speed int
}

type ferrari struct {
	car
	turbo bool
}

func main() {
	f := ferrari{}
	f.name = "F40"
	f.speed = 0
	f.turbo = true
}
