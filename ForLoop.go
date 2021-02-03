package main

import "fmt"

func ForLoop() {

	stringSlice := []string{
		"firstString",
		"secondString",
		"thirdString",
		"fourthString",
	}

	for _, string := range stringSlice {
		fmt.Println(string)
	}

}
