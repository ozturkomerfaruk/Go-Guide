package main

import "fmt"

func maind() {
	v1, v2 := split(5)
	fmt.Println(v1, v2)

	count, sum := add(3,2,1)
	fmt.Println(count, sum)

	//Anonim Fonksiyonları var birde
	addFunc := func(terms ...int) (numTerms int, sum int) {
		for _, term := range terms {
			sum += term
		}
		numTerms = len(terms)
		return
	}

	numTerms, sum := addFunc(5,2)
	println(numTerms, sum)
}

func split(sum float32) (x, y float32) {
	x = sum * 2
	y = sum / 2
	return
}

func add(terms ...int) (numTerms int, sum int) {
	for _, term := range terms {
		sum += term
	}
	numTerms = len(terms)
	return
}