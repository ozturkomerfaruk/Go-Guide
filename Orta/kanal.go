package main

import "fmt"

func main() {
    // Bir meyve kanalı oluştur
    fruitChannel := make(chan string)

    // Elma kanala gönder
    go func() {
        fruitChannel <- "Elma"
    }()

    // Armut kanala gönder
    go func() {
        fruitChannel <- "Armut"
    }()

    // Portakal kanala gönder
    go func() {
        fruitChannel <- "Portakal"
    }()

    // Kanaldan meyve al
    fruit1 := <-fruitChannel
    fruit2 := <-fruitChannel
    fruit3 := <-fruitChannel

    // Alınan meyveleri yazdır
    fmt.Println(fruit1, fruit2, fruit3)
}