package main

import (
	"errors"
	"fmt"
)

func mainb() {
	//if let
	if foo := 2; foo == 2 {
		fmt.Println("koşulda")
	} else {
		fmt.Println("koşulda değil")
	}

	//array
	array := [...]string{"a", "b", "c", "d"}
	for i := range array {
		fmt.Println("Array item: ", i, " is", array[i])
	}

	//map
	baskentler := map[string]string{"Turkey": "Ankara", "France": "Paris"}
	for key := range baskentler {
		fmt.Println("Map item: Capital of", key, "is", baskentler[key])
	}

	//error
	myError := errors.New("Bu bir hata!")
	fmt.Println(myError)

	//slice
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	//slice örnek 2
	myArr := [...]int{999, 3, 23, 4, 23, 2312, 4}
	mySlice1 := myArr[:]
	fmt.Println(mySlice1)
	mySlice1 = append(mySlice1, 0)
	mySlice2 := append(mySlice1[1:len(myArr)])
	fmt.Println(mySlice2, "asd")

	/*
		Go dilinde, len() ve cap() fonksiyonları bir dizi veya slice'in farklı özelliklerini döndürür.
			len() fonksiyonu, bir dizi veya slice'in eleman sayısını döndürür.
			cap() fonksiyonu, bir slice'in altında yatan dizi kapasitesini döndürür.
		Dizilerin boyutu, dizi tanımlanırken belirlendiğinden, bir dizi için kapasite kavramı yoktur.
		Bu nedenle, cap() fonksiyonu yalnızca slice'larla kullanılabilir.
	*/

	bigArr := make([]int, 5)
	bigArr = append(bigArr, 1)

	fmt.Println(len(bigArr), "len bigArr")
	fmt.Println(cap(bigArr), "cap bigArr")
	//cap 10 değeri döndürür ama len 6 değerini döndürür. Bunun en büyük sebebi cap de kapasite artırıldı
	//ve bu 5lik olarak artırıldı.
}
