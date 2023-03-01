package main

import (
	"fmt"
	"math/rand"
	"strings"

	"omer.com/faruk/utils"
)

func mainc() {
	//Map
	myMap := make(map[int]string)
	myMap[10] = "on"
	myMap[6] = "altı"
	fmt.Println(myMap)

	delete(myMap, 6)
	fmt.Println(myMap)

	//Packages
	fmt.Println("number is:", rand.Intn(10))

	//String paket
	fmt.Println(strings.Contains("test", "es"))
	fmt.Println(strings.Count("test", "t"))
	fmt.Println(strings.HasPrefix("unit_test", "unit"))
	fmt.Println(strings.Index("test", "e"))

	n1, l1 := utils.FullName("Ömer Faruk", "Öztürk")
	fmt.Printf("Full name: %v, number of chars: %v\n\n", n1, l1)
}
