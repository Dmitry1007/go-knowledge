package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

type boat struct {
	Make  string
	Model string
	Year  int
}

func main() {
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}

	os.Stdout.Write(b)

	var cg ColorGroup
	error := json.Unmarshal(b, &cg)
	if error != nil {
		fmt.Println("error:", error)
	}

	fmt.Printf("\n %+v is of Type %T \n", cg, cg)

	yat := boat{Make: "Amel", Model: "Maramu", Year: 2020}
	cat := boat{Make: "Lagoon", Model: "50", Year: 2021}
	boats := []boat{yat, cat}
	fmt.Println("Boats", boats)

	j, err := json.Marshal(boats)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Marshaled to json", string(j))

	var yatchs []boat
	err = json.Unmarshal([]byte(j), &yatchs)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Unmarshaled from Json", yatchs)

	for i, v := range yatchs {
		fmt.Println("yatch number", i)
		fmt.Println("\t", v.Make, v.Model, v.Year)
	}
}
