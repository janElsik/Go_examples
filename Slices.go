package main

import "fmt"

func Slices() {

	//slice of strings
	stringSlice := []string{
		"firstString",
		"secondString",
		"thirdString",
		"fourthString",
	}

	for _, string := range stringSlice {
		fmt.Println(string)
	}

	fmt.Println("======================================")

	// slice of int
	var i int = 16
	var j int = 5
	var k int = i + j

	intSlice := []int{i, j, k, 20}

	var m int = 125

	// adding another element to the slice
	intSlice = append(intSlice, m)

	for _, int := range intSlice {
		fmt.Println(int)
	}

	names:= [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a,b)

	b[0] = "XXX"
	fmt.Println(a,b)

	fmt.Println(names)

	// slice literals

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)


	//Slice defaults - slicing piece by piece
	t := []int{2, 3, 5, 7, 11, 13}

	t= t[1:4]
	fmt.Println(s)

	t = t[:2]
	fmt.Println(s)



	t = t[1:]
	fmt.Println(s)
}
