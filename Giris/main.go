package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func main() {
	// String to Int
	var myString = "1"
	number, err := strconv.Atoi(myString) // err yerine _ koyabilirsin.
	fmt.Println(number, err)

	// Int to String
	myInt := 2
	stringNum := strconv.Itoa(myInt)
	fmt.Println(stringNum)

	//Output
	var mystr string = "deneme deneme alo alo"
	strength, _ := fmt.Println(mystr)
	fmt.Println(strength)

	isTrue := true
	fmt.Printf("Value of isTrue: %v\n", isTrue)

	fmt.Printf("Data Type -> %s\n", reflect.TypeOf(isTrue))
	fmt.Printf("Data Type -> %T\n", isTrue)

	//Input
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter a num: ")
	str, _ := reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Sayi -> %v\n", f)
	}

	//Time
	nowTime := time.Now()
	fmt.Println(nowTime)
}
