package main

import "fmt"

type Human struct {
	Firstname string
	Lastname  string
	Age       int
}

func NewHuman() *Human {
	h := new(Human)
	return h
}

func NewHumanParam(firstname, lastname string, age int) *Human {
	h := new(Human)
	h.Firstname = firstname
	h.Lastname = lastname
	h.Age = age
	return h
}

func mainw() {
	x := NewHumanParam("ali", "veli", 50)
	fmt.Println(x)
}
