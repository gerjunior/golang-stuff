package main

import "fmt"

func salariesFunc() {
	salaries := map[string]float64{
		"John":  10000.34,
		"Jane":  20000.00,
		"Marko": 1200.00,
	}

	salaries["John"] = 15000.00
	delete(salaries, "Unknown")

	for name, salary := range salaries {
		fmt.Println(name, salary)
	}

	fmt.Println(salaries)
}

func mapInsideMap() {
	funcs := map[string]map[string]float64{
		"G": {
			"Gabriel": 3321,
			"Geraldo": 1234,
		},
		"J": {
			"John": 1313,
			"Jane": 3211,
		},
	}

	for l, list := range funcs {
		for name, salary := range list {
			fmt.Printf("%v (%v) => %v\n", name, l, salary)
		}
	}
}

func mapToTheInfinity() {
	mapOfMappers := make(map[string]interface{})

	mapOfMappers["A"] = map[string]interface{}{
		"B": map[string]interface{}{
			"C": map[string]string{
				"D": "A",
			},
		},
	}
	fmt.Println(mapOfMappers)
}

func main() {
	// approved := make(map[int]string)

	// approved[12312312] = "John"
	// approved[12312313] = "Jane"
	// approved[12312314] = "None"

	// for cpf, name := range approved {
	// 	fmt.Printf("%v (CPF: %v)\n", name, cpf)
	// }

	// delete(approved, 12312314)

	// salariesFunc()
	// mapInsideMap()
	mapToTheInfinity()
}
