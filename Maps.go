package main

import (
	"fmt"
)

type Vvertex struct {
	Lat, Long float64
}

var map1 = map[string]Vvertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(map1)

	map2 := make(map[string]int)

	map2["Answer"] = 42
	fmt.Println("The value:", map2["Answer"])

	map2["Answer"] = 48
	fmt.Println("The value:", map2["Answer"])

	delete(map2, "Answer")
	fmt.Println("The value:", map2["Answer"])

	v, ok := map2["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
