package main

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomNumber() {

	var i int = 1

	fmt.Println("Value of int i is", i)

	//declare a seed for the system to look into
	rand.Seed(time.Now().UTC().UnixNano())

	// rand.Intn() uses the random seed
	// if the random seed would not be there,
	// than the value of j would be infinitely same
	j := rand.Intn(20)

	fmt.Println(j)

}
