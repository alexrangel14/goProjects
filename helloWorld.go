// 3:21:17
// https://www.youtube.com/watch?v=YS4e4q9oBaU

// go build fileName.go
// ./fileName

package main

import (
	"fmt"
	"strconv"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

func main() {
	////////////////////////////////////////////////////////////////////////
	// Struct
	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
	fmt.Println(aDoctor.companions)

	////////////////////////////////////////////////////////////////////////
	fmt.Println("Hello!")
	////////////////////////////////////////////////////////////////////////

	// Variables

	// One way to declare a var
	var firstWay int = 42
	fmt.Println(firstWay) // prints 42

	// Another way to declare a var
	secondWay := 63.89
	fmt.Println(secondWay)

	//Explicit conversion
	i := 90
	j := int(i)
	fmt.Println(j)

	//convert int to string
	k := strconv.Itoa(i)
	fmt.Println(k)

	////////////////////////////////////////////////////////////////////////
	// Primitives

	// Boolean
	boolVar := true
	fmt.Println(boolVar)

	// Int
	number := 42
	number2 := 21
	fmt.Println(number)

	// Math Operations
	fmt.Println(number + number2)
	fmt.Println(number - number2)
	fmt.Println(number / number2)
	fmt.Println(number % number2)

	// Strings
	s := "this is a string"
	s2 := "this is also a string"
	fmt.Println(s)

	// Get a character from string
	fmt.Println(string(s[2]))

	// Add strings together
	fmt.Println(s + s2)
	////////////////////////////////////////////////////////////////////////
	// Constants
	const myConst int = 42
	fmt.Println(myConst)
	////////////////////////////////////////////////////////////////////////
	// Arrays and slices

	// arrayName := [arraySize]arrayType{intializerValues}
	grades := [3]int{97, 85, 93}
	fmt.Printf("Grades: %v \n", grades)

	fmt.Println(grades[2])
	fmt.Println(len(grades))

	// Slice
	c := []int{}
	fmt.Println(c)
	fmt.Println(len(c))
	fmt.Println(cap(c))
	c = append(c, 1)
	fmt.Println(c)
	fmt.Println(len(c))
	fmt.Println(cap(c))
	c = append(c, 5, 6, 7, 8)
	fmt.Println(c)
	fmt.Println(len(c))
	fmt.Println(cap(c))
	////////////////////////////////////////////////////////////////////////
	// Maps and Structs
	//map[key]value
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      201612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	// Add to the map
	statePopulations["Georgia"] = 10310371
	// Remove from the map
	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations)
	// "ok" will be bool determing if key is in map
	_, ok := statePopulations["Oho"]
	fmt.Println(ok)
	////////////////////////////////////////////////////////////////////////
	// If and switch statements
	if false {
		fmt.Println("The test is true")
	}

	if pop, ok := statePopulations["Florida"]; ok {
		fmt.Println(pop)
	}

	numberNew := 50
	guess := 50

	if guess < 1 || guess > 100 {
		fmt.Println("The guess must be between 1 and 100")
	} else if guess < 1 || guess > 100 {
		fmt.Println("Number must be in range (1 - 100)")
	} else {
		if guess < numberNew {
			fmt.Println("Too low")
		}
		if guess > numberNew {
			fmt.Println("Too high")
		}
		if guess == numberNew {
			fmt.Println("You got it!")
		}
	}
}

func returnTrue() bool {
	fmt.Println("Returning true")
	return true
}
